package handlers

import (
	"bitbucket.com/finease/backend/pkg/controllers/api"
	"bitbucket.com/finease/backend/pkg/controllers/services"
	"bitbucket.com/finease/backend/pkg/models"
	"bitbucket.com/finease/backend/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoanProposal interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type loanProposalsHandler struct {
	loanProposalsService services.LoanProposal
}

func NewLoanProposalsHandler(loanProposalsService services.LoanProposal) LoanProposal {
	return &loanProposalsHandler{loanProposalsService: loanProposalsService}
}

func (l loanProposalsHandler) Create(c *gin.Context) {
	var reqBody api.LoanProposal
	if err := c.BindJSON(&reqBody); err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to parse the request body: %v", err))
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	user, ok := c.Keys["user"].(models.User)
	if !ok {
		resp := utils.ResponseRenderer("failed to parse the user uuid from the processed header")
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	userUuid := user.Uuid

	inboundLoanProposalModel := api.MapLoanProposalRequestToModel(&reqBody)
	inboundLoanProposalModel.UserUUID = userUuid

	createdLoanProposal, err := l.loanProposalsService.Create(c, inboundLoanProposalModel)
	if err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to create the loan proposal: %v", err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	outboundLoanProposal := api.MapLoanProposalModelToResponse(createdLoanProposal)

	resp := utils.ResponseRenderer("Loan Proposal created successfully", gin.H{
		"loan_proposal": outboundLoanProposal,
	})
	c.JSON(http.StatusOK, resp)
}

func (l loanProposalsHandler) Delete(c *gin.Context) {
	loanProposalUuid := c.Param("loan_proposal_uuid")
	user, ok := c.Keys["user"].(models.User)
	if !ok {
		resp := utils.ResponseRenderer("failed to parse the user uuid from the processed header")
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	userUuid := user.Uuid

	if err := l.loanProposalsService.Delete(c, userUuid, loanProposalUuid); err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to delete the loan proposal: %v", err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := utils.ResponseRenderer("Loan proposal deleted successfully")
	c.JSON(http.StatusOK, resp)
}

func (l loanProposalsHandler) List(c *gin.Context) {
	user, ok := c.Keys["user"].(models.User)
	if !ok {
		resp := utils.ResponseRenderer("failed to parse the user uuid from the processed header")
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	userUuid := user.Uuid
	loanProposals, err := l.loanProposalsService.Find(c, userUuid)
	if err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to list your loan proposals: %w", err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	outboundLoanProposals := make([]*api.LoanProposal, len(loanProposals))
	for i, loanProposal := range loanProposals {
		outboundLoanProposals[i] = api.MapLoanProposalModelToResponse(loanProposal)
	}
	resp := utils.ResponseRenderer("Your loan proposals fetched successfully", gin.H{
		"loan_proposals": outboundLoanProposals,
	})
	c.JSON(http.StatusOK, resp)

}

func (l loanProposalsHandler) Update(c *gin.Context) {
	loanProposalUuid := c.Param("loan_proposal_uuid")
	user, ok := c.Keys["user"].(models.User)
	if !ok {
		resp := utils.ResponseRenderer("failed to parse the user uuid from the processed header")
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	userUuid := user.Uuid

	var reqBody api.LoanProposal
	if err := c.BindJSON(&reqBody); err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to parse the request body: %v", err))
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	inboundLoanProposalModel := api.MapLoanProposalRequestToModel(&reqBody)
	updatedLoanProposal, err := l.loanProposalsService.Update(c, userUuid, loanProposalUuid, inboundLoanProposalModel)
	if err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to update the loan proposal: %v", err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	outboundLoanProposal := api.MapLoanProposalModelToResponse(updatedLoanProposal)

	resp := utils.ResponseRenderer("Loan Proposal updated successfully", gin.H{
		"loan_proposal": outboundLoanProposal,
	})
	c.JSON(http.StatusOK, resp)
}
