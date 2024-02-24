package v1

import (
	"bitbucket.com/finease/backend/pkg/controllers/handlers"
	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(parentRouter *gin.RouterGroup, authHandler handlers.Auth, additionalMiddlewares ...gin.HandlerFunc) {
	authRouter := parentRouter.Group("/auth")

	authRouter.Use(additionalMiddlewares...)

	authRouter.POST("/register", authHandler.Register)
	authRouter.POST("/login", authHandler.Login)
}
