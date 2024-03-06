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
	FindByUserId(c *gin.Context)
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

	inboundLoanRequestModel := api.MapLoanRequestApiToModel(&reqBody)
	inboundLoanRequestModel.UserUUID = userUuid

	createdLoanRequest, err := l.loanRequestService.Create(c, inboundLoanRequestModel)
	if err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to create the loan request: %v", err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	outboundLoanRequest := api.MapLoanRequestModelToApi(createdLoanRequest)

	resp := utils.ResponseRenderer("Loan Request created successfully", gin.H{
		"loan_request": outboundLoanRequest,
	})
	c.JSON(http.StatusOK, resp)
}

func (l loanRequestHandler) Delete(c *gin.Context) {
	loanRequestUuid := c.Param("loan_request_uuid")
	user, ok := c.Keys["user"].(models.User)
	if !ok {
		resp := utils.ResponseRenderer("failed to parse the user uuid from the processed header")
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	userUuid := user.Uuid

	if err := l.loanRequestService.Delete(c, userUuid, loanRequestUuid); err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to delete the loan request: %v", err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := utils.ResponseRenderer("Loan request deleted successfully")
	c.JSON(http.StatusOK, resp)
}

func (l loanRequestHandler) Update(c *gin.Context) {
	loanRequestUuid := c.Param("loan_request_uuid")
	user, ok := c.Keys["user"].(models.User)
	if !ok {
		resp := utils.ResponseRenderer("failed to parse the user uuid from the processed header")
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	userUuid := user.Uuid

	var reqBody api.LoanRequest
	if err := c.BindJSON(&reqBody); err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to parse the request body: %v", err))
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	inboundLoanRequestModel := api.MapLoanRequestApiToModel(&reqBody)
	updatedLoanrequest, err := l.loanRequestService.Update(c, userUuid, loanRequestUuid, inboundLoanRequestModel)
	if err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to update the loan request: %v", err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	outboundLoanrequest := api.MapLoanRequestModelToApi(updatedLoanrequest)

	resp := utils.ResponseRenderer("Loan request updated successfully", gin.H{
		"loan_request": outboundLoanrequest,
	})
	c.JSON(http.StatusOK, resp)
}

func (l loanRequestHandler) FindByUserId(c *gin.Context) {
	user, ok := c.Keys["user"].(models.User)
	if !ok {
		resp := utils.ResponseRenderer("failed to parse the user uuid from the processed header")
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	userUuid := user.Uuid
	loanRequests, err := l.loanRequestService.FindByUserId(c, userUuid)

	if err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to list your loan Request: %v", err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	outboundLoanRequest := make([]*api.LoanRequest, len(loanRequests))
	for i, loanRequest := range loanRequests {
		outboundLoanRequest[i] = api.MapLoanRequestModelToApi(loanRequest)
	}
	resp := utils.ResponseRenderer("Your loan Request fetched successfully", gin.H{
		"loan_Requests": outboundLoanRequest,
	})
	c.JSON(http.StatusOK, resp)
}

func (l loanRequestHandler) FindById(c *gin.Context) {
	loanRequest, ok := c.Keys["loan_request"].(models.LoanRequest)
	if !ok {
		resp := utils.ResponseRenderer("failed to parse the user uuid from the processed header")
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	loanRequestUuid := loanRequest.Uuid
	loanRequestResp, err := l.loanRequestService.FindById(c, loanRequestUuid)

	if err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to list your loan Request: %v", err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	outboundLoanRequest := api.MapLoanRequestModelToApi(loanRequestResp)

	resp := utils.ResponseRenderer("Your loan Request fetched successfully", gin.H{
		"loan_Request": outboundLoanRequest,
	})
	c.JSON(http.StatusOK, resp)
}
