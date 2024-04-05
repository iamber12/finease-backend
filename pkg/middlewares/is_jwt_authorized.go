package middlewares

import (
	"fmt"
	"net/http"

	"bitbucket.com/finease/backend/pkg/dao"
	"bitbucket.com/finease/backend/pkg/environment"
	"bitbucket.com/finease/backend/pkg/utils"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func IsJwtAuthorized(userDao dao.User) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		applicationConfig := environment.Env.ApplicationConfig
		jwtSecret := applicationConfig.ServerConfig.JwtSecret
		if jwtSecret == "" {
			jwtSecret = "success"
		}

		xAccessToken := c.GetHeader("X-Access-Token")
		if xAccessToken == "" {
			resp := utils.ResponseRenderer("Token not found!")
			c.AbortWithStatusJSON(http.StatusBadRequest, resp)
			return
		}
		claims := &utils.JwtClaims{}
		token, err := jwt.ParseWithClaims(xAccessToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil {
			code := http.StatusBadRequest
			if err == jwt.ErrSignatureInvalid {
				code = http.StatusUnauthorized
			}
			resp := utils.ResponseRenderer(fmt.Errorf("Invalid JWT found: %w", err).Error())
			c.AbortWithStatusJSON(code, resp)
			return
		}
		if !token.Valid {
			resp := utils.ResponseRenderer("Invalid JWT found!")
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}

		if len(c.Keys) == 0 {
			c.Keys = map[string]any{}
		}

		userUuid := claims.UserUuid
		foundUser, err := userDao.FindById(c, userUuid)

		if err != nil {
			resp := utils.ResponseRenderer("The token found to belong to a non-existent user")
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}

		if !utils.FromPtr(foundUser.Active) {
			resp := utils.ResponseRenderer("The user has been deactivated")
			c.AbortWithStatusJSON(http.StatusForbidden, resp)
			return
		}

		c.Keys["user"] = *foundUser
		c.Next()
	})
}
