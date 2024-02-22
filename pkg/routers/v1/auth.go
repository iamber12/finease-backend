package v1

import (
	"bitbucket.com/finease/backend/pkg/controllers/auth/login"
	"bitbucket.com/finease/backend/pkg/controllers/auth/register"
	"bitbucket.com/finease/backend/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(parentRouter *gin.RouterGroup, additionalMiddlewares ...gin.HandlerFunc) {
	authRouter := parentRouter.Group("/auth")

	authRouter.Use(additionalMiddlewares...)

	authRouter.GET("/login", login.Controller)
	authRouter.GET("/register", middlewares.IsJwtAuthorized, register.Controller)
}
