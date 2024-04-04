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
	loanRequestsRouter.PUT("/request/:loan_request_uuid", loanRequestsHandler.Update)
	loanRequestsRouter.DELETE("/request/:loan_request_uuid", loanRequestsHandler.Delete)
	loanRequestsRouter.GET("/available", loanRequestsHandler.FindAll)
	loanRequestsRouter.GET("/request/:loan_request_uuid", loanRequestsHandler.FindById)

	loanRequestsRouter.PUT("/request/:loan_request_uuid/accept", loanRequestsHandler.Accept) // to be executed by the lender
	loanRequestsRouter.PUT("/request/:loan_request_uuid/reject", loanRequestsHandler.Reject) // to be executed by the lender
	loanRequestsRouter.GET("/received", loanRequestsHandler.GetReceivedForProposal)          // to be executed by the lender
	loanRequestsRouter.GET("/requested", loanRequestsHandler.GetRequestedForUser)            // to be executed by the borrower
}
