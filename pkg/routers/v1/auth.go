package v1

import (
	"bitbucket.com/finease/backend/pkg/controllers/handlers"
	"bitbucket.com/finease/backend/pkg/controllers/services"
	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(parentRouter *gin.RouterGroup, authService services.Auth) {
	authRouter := parentRouter.Group("/auth")

	authHandler := handlers.NewAuthHandler(authService)

	authRouter.POST("/register", authHandler.Register)
	authRouter.POST("/login", authHandler.Login)
}
