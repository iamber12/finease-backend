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
	loanProposalRouter.GET("/my", loanProposalHandler.ListMine)
	loanProposalRouter.GET("/available", loanProposalHandler.ListAvailable)
	loanProposalRouter.DELETE("/:loan_proposal_uuid", loanProposalHandler.Delete)

	loanProposalGrantRouter := loanProposalRouter.Group("/:loan_proposal_uuid/grant")
	loanProposalGrantRouter.PUT("/offer", loanProposalHandler.OfferGrant)   // to be executed by the lender
	loanProposalGrantRouter.PUT("/revoke", loanProposalHandler.RevokeGrant) // to be executed by the lender

	loanProposalGrantRouter.PUT("/accept", loanProposalHandler.AcceptGrant) // to be executed by the borrower
	loanProposalGrantRouter.PUT("/reject", loanProposalHandler.RejectGrant) // to be executed by the borrower

}
