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

type FinancialTransaction interface {
	Create(c *gin.Context)
	FindAll(c *gin.Context)
	FindReceived(c *gin.Context)
	FindSent(c *gin.Context)
	FindByLoanProposalUuid(c *gin.Context)
	FindByLoanRequestUuid(c *gin.Context)
}

type financialTransactionHandler struct {
	financialTransactionService services.FinancialTransaction
}

func NewFinancialTransactionHandler(financialTransactionService services.FinancialTransaction) FinancialTransaction {
	return &financialTransactionHandler{financialTransactionService: financialTransactionService}
}

func (f financialTransactionHandler) Create(c *gin.Context) {
	var reqBody api.FinancialTransaction
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

	inboundTransactionModel := api.MapFinancialTransactionApiToModel(&reqBody)
	if inboundTransactionModel.LenderUuid != userUuid && inboundTransactionModel.BorrowerUuid != userUuid {
		resp := utils.ResponseRenderer("cannot create financial " +
			"transaction in which you aren't the borrower/lender")
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	createdTransaction, err := f.financialTransactionService.Create(c, inboundTransactionModel)
	if err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to create the transaction: %v", err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	outboundTransaction := api.MapFinancialTransactionModelToApi(createdTransaction)

	resp := utils.ResponseRenderer("Financial Transaction created successfully", gin.H{
		"transaction": outboundTransaction,
	})
	c.JSON(http.StatusOK, resp)
}

func (f financialTransactionHandler) FindAll(c *gin.Context) {
	user, ok := c.Keys["user"].(models.User)
	if !ok {
		resp := utils.ResponseRenderer("failed to parse the user uuid from the processed header")
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	transactions, err := f.financialTransactionService.FindAll(c, user.Uuid)
	if err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to list your transactions: %v", err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	outboundTransactions := make([]*api.FinancialTransaction, len(transactions))
	for i, transaction := range transactions {
		outboundTransactions[i] = api.MapFinancialTransactionModelToApi(transaction)
	}
	resp := utils.ResponseRenderer("Your transactions fetched successfully", gin.H{
		"transactions": outboundTransactions,
	})
	c.JSON(http.StatusOK, resp)
}

func (f financialTransactionHandler) FindReceived(c *gin.Context) {
	user, ok := c.Keys["user"].(models.User)
	if !ok {
		resp := utils.ResponseRenderer("failed to parse the user uuid from the processed header")
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	transactions, err := f.financialTransactionService.FindReceived(c, user.Uuid)
	if err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to list your transactions: %v", err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	outboundTransactions := make([]*api.FinancialTransaction, len(transactions))
	for i, transaction := range transactions {
		outboundTransactions[i] = api.MapFinancialTransactionModelToApi(transaction)
	}
	resp := utils.ResponseRenderer("The transactions you received fetched successfully", gin.H{
		"transactions": outboundTransactions,
	})
	c.JSON(http.StatusOK, resp)
}

func (f financialTransactionHandler) FindSent(c *gin.Context) {
	user, ok := c.Keys["user"].(models.User)
	if !ok {
		resp := utils.ResponseRenderer("failed to parse the user uuid from the processed header")
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	transactions, err := f.financialTransactionService.FindSent(c, user.Uuid)
	if err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to list your transactions: %v", err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	outboundTransactions := make([]*api.FinancialTransaction, len(transactions))
	for i, transaction := range transactions {
		outboundTransactions[i] = api.MapFinancialTransactionModelToApi(transaction)
	}
	resp := utils.ResponseRenderer("The transactions you sent fetched successfully", gin.H{
		"transactions": outboundTransactions,
	})
	c.JSON(http.StatusOK, resp)
}

func (f financialTransactionHandler) FindByLoanProposalUuid(c *gin.Context) {
	user, ok := c.Keys["user"].(models.User)
	if !ok {
		resp := utils.ResponseRenderer("failed to parse the user uuid from the processed header")
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	proposalUuid := c.Param("loan_proposal_uuid")

	transactions, err := f.financialTransactionService.FindByLoanProposalUuid(c, user.Uuid, proposalUuid)
	if err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to list your transactions: %v", err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	outboundTransactions := make([]*api.FinancialTransaction, len(transactions))
	for i, transaction := range transactions {
		outboundTransactions[i] = api.MapFinancialTransactionModelToApi(transaction)
	}
	resp := utils.ResponseRenderer("The transactions associated with the proposal fetched successfully", gin.H{
		"transactions": outboundTransactions,
	})
	c.JSON(http.StatusOK, resp)
}

func (f financialTransactionHandler) FindByLoanRequestUuid(c *gin.Context) {
	user, ok := c.Keys["user"].(models.User)
	if !ok {
		resp := utils.ResponseRenderer("failed to parse the user uuid from the processed header")
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	requestUuid := c.Param("loan_request_uuid")

	transactions, err := f.financialTransactionService.FindByLoanRequestUuid(c, user.Uuid, requestUuid)
	if err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to list your transactions: %v", err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	outboundTransactions := make([]*api.FinancialTransaction, len(transactions))
	for i, transaction := range transactions {
		outboundTransactions[i] = api.MapFinancialTransactionModelToApi(transaction)
	}
	resp := utils.ResponseRenderer("The transactions associated with the loan request fetched successfully", gin.H{
		"transactions": outboundTransactions,
	})
	c.JSON(http.StatusOK, resp)
}
