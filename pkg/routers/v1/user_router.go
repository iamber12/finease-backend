package v1

import (
	"bitbucket.com/finease/backend/pkg/controllers/handlers"
	"bitbucket.com/finease/backend/pkg/controllers/services"
	"github.com/gin-gonic/gin"
)

func SetupUserRouter(parentRouter *gin.RouterGroup, userService services.User, additionalMiddlewares ...gin.HandlerFunc) {
	userRouter := parentRouter.Group("/user")
	userHandler := handlers.NewUserHandler(userService)

	userRouter.Use(additionalMiddlewares...)

	userRouter.GET("/", userHandler.FindById)
	userRouter.GET("/:user_uuid", userHandler.FindById)
	userRouter.PUT("/", userHandler.Update)
	userRouter.PUT("/deactivate", userHandler.Deactivate)
}
