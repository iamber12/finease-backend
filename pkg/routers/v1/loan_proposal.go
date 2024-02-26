package v1

import (
	"bitbucket.com/finease/backend/pkg/controllers/handlers"
	"bitbucket.com/finease/backend/pkg/controllers/services"
	"github.com/gin-gonic/gin"
)

func SetupLoanProposalRouter(parentRouter *gin.RouterGroup, loanProposalService services.LoanProposal, additionalMiddlewares ...gin.HandlerFunc) {
	loanProposalRouter := parentRouter.Group("/loan/proposals")
	loanProposalHandler := handlers.NewLoanProposalsHandler(loanProposalService)

	loanProposalRouter.Use(additionalMiddlewares...)

	loanProposalRouter.POST("/", loanProposalHandler.Create)
	loanProposalRouter.PUT("/:loan_proposal_uuid", loanProposalHandler.Update)
	loanProposalRouter.GET("/", loanProposalHandler.List)
	loanProposalRouter.DELETE("/:loan_proposal_uuid", loanProposalHandler.Delete)
}
