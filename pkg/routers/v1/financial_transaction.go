package v1

import (
	"bitbucket.com/finease/backend/pkg/controllers/handlers"
	"bitbucket.com/finease/backend/pkg/controllers/services"
	"github.com/gin-gonic/gin"
)

func SetupFinancialTransactionRouter(parentRouter *gin.RouterGroup, financialTransactionService services.FinancialTransaction, additionalMiddlewares ...gin.HandlerFunc) {
	financialTransactionRouter := parentRouter.Group("/transactions")
	financialTransactionHandler := handlers.NewFinancialTransactionHandler(financialTransactionService)

	financialTransactionRouter.Use(additionalMiddlewares...)

	financialTransactionRouter.POST("/", financialTransactionHandler.Create)
	financialTransactionRouter.GET("/", financialTransactionHandler.FindAll)
	financialTransactionRouter.GET("/received", financialTransactionHandler.FindReceived)
	financialTransactionRouter.GET("/sent", financialTransactionHandler.FindSent)
	financialTransactionRouter.GET("/proposal/:loan_proposal_uuid", financialTransactionHandler.FindByLoanProposalUuid)
	financialTransactionRouter.GET("/request/:loan_request_uuid", financialTransactionHandler.FindByLoanRequestUuid)
}
