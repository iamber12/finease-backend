package v1

import (
	"bitbucket.com/finease/backend/pkg/controllers/handlers"
	"bitbucket.com/finease/backend/pkg/controllers/services"
	"github.com/gin-gonic/gin"
)

func SetupLoanRequestsRouter(parentRouter *gin.RouterGroup, loanRequestsService services.LoanRequest, additionalMiddlewares ...gin.HandlerFunc) {
	loanRequestsRouter := parentRouter.Group("/loan/requests")
	loanRequestsHandler := handlers.NewLoanRequestHandler(loanRequestsService)

	loanRequestsRouter.Use(additionalMiddlewares...)

	loanRequestsRouter.POST("/", loanRequestsHandler.Create)
	loanRequestsRouter.PUT("/:loan_request_uuid", loanRequestsHandler.Update)
	loanRequestsRouter.DELETE("/:loan_request_uuid", loanRequestsHandler.Delete)
	loanRequestsRouter.GET("/", loanRequestsHandler.FindByUserId)
	loanRequestsRouter.GET("/all", loanRequestsHandler.FindAll)
	loanRequestsRouter.GET("/:loan_request_uuid", loanRequestsHandler.FindById)
	loanRequestsRouter.GET("/proposal/:loan_proposal_uuid", loanRequestsHandler.FindByProposalId)
}
