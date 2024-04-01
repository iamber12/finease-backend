package v1

import (
	"bitbucket.com/finease/backend/pkg/controllers/handlers"
	"bitbucket.com/finease/backend/pkg/controllers/services"
	"github.com/gin-gonic/gin"
)

func SetupUserRouter(parentRouter *gin.RouterGroup, userService services.User, additionalMiddlewares ...gin.HandlerFunc) {
	loanProposalRouter := parentRouter.Group("/user")
	loanProposalHandler := handlers.NewUserHandler(userService)

	loanProposalRouter.Use(additionalMiddlewares...)

	loanProposalRouter.GET("/", loanProposalHandler.FindById)
	loanProposalRouter.GET("/:user_uuid", loanProposalHandler.FindById)
	loanProposalRouter.PUT("/", loanProposalHandler.Update)
	loanProposalRouter.PUT("/deactivate", loanProposalHandler.Deactivate)
}
