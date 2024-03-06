package v1

import (
	"bitbucket.com/finease/backend/pkg/controllers/handlers"
	"github.com/gin-gonic/gin"
)

func SetupLoanRequestsRouter(parentRouter *gin.RouterGroup, loanRequestsHandler handlers.LoanRequest, additionalMiddlewares ...gin.HandlerFunc) {
	authRouter := parentRouter.Group("/loan/request")

	authRouter.Use(additionalMiddlewares...)

	authRouter.POST("/", loanRequestsHandler.Create)
	authRouter.PUT("/", loanRequestsHandler.Update)
	authRouter.DELETE("/", loanRequestsHandler.Delete)
	authRouter.GET("/", loanRequestsHandler.FindByUserId)
	authRouter.GET("/", loanRequestsHandler.FindById)
}
