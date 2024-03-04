package handlers

import (
	"fmt"
	"net/http"

	"bitbucket.com/finease/backend/pkg/controllers/api"
	"bitbucket.com/finease/backend/pkg/controllers/services"
	"bitbucket.com/finease/backend/pkg/models"
	"bitbucket.com/finease/backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

type LoanRequest interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
}

type loanRequestHandler struct {
	loanRequestService services.LoanRequest
}

func NewLoanRequestHandler(loanRequestService services.LoanRequest) LoanRequest {
	return &loanRequestHandler{loanRequestService: loanRequestService}
}

func (l loanRequestHandler) Create(c *gin.Context) {
	var reqBody api.LoanRequest

	if err := c.BindJSON; err != nil {
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

	inboundLoanRequestModel := api.MapLoanRequestApiToModel(&reqBody)
	inboundLoanRequestModel.UserUUID = userUuid

	createdLoanRequest, err := l.loanRequestService.Create(c, inboundLoanRequestModel)
	if err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to create the loan proposal: %v", err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	outboundLoanRequest := api.MapLoanRequestModelToApi(createdLoanRequest)

	resp := utils.ResponseRenderer("Loan Proposal created successfully", gin.H{
		"loan_proposal": outboundLoanRequest,
	})
	c.JSON(http.StatusOK, resp)
}
