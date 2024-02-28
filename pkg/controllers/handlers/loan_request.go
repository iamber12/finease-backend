package handlers

import (
	"bitbucket.com/finease/backend/pkg/controllers/services"
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
	authService services.Auth
}

// func NewLoanRequestHandler(authService services.Auth) Auth {
// 	return &loanRequestHandler{authService: authService}
// }
