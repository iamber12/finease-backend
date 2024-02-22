package middlewares

import (
	"bitbucket.com/finease/backend/pkg/environment"
	"bitbucket.com/finease/backend/pkg/utils"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JwtClaims struct {
	UserUUID string `json:"user_uuid,omitempty"`
	jwt.StandardClaims
}

func GenerateJWT(userUuid, secret string, expiresAt int64) (string, error) {
	claims := &JwtClaims{
		UserUUID: userUuid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func IsJwtAuthorized(c *gin.Context) {
	applicationConfig := environment.Env.ApplicationConfig
	jwtSecret := applicationConfig.ServerConfig.JwtSecret

	xAccessToken := c.GetHeader("X-Access-Token")
	if xAccessToken == "" {
		resp := utils.ResponseRenderer("Token not found!")
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}
	claims := &JwtClaims{}
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
	c.Keys["user_uuid"] = claims.UserUUID
	c.Next()
}
