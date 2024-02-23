package handlers

import (
	"bitbucket.com/finease/backend/pkg/controllers/services"
	"bitbucket.com/finease/backend/pkg/models"
	"bitbucket.com/finease/backend/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Auth interface {
	Register(c *gin.Context)
}

type authHandler struct {
	authService services.Auth
}

func NewAuthHandler(authService services.Auth) Auth {
	return &authHandler{authService: authService}
}

func (h authHandler) Register(c *gin.Context) {
	type RequestBody struct {
		Name        string `json:"name,omitempty"`
		DateOfBirth string `json:"date_of_birth,omitempty"`
		Address     string `json:"address,omitempty"`
		PrimaryRole string `json:"primary_role,omitempty"`
		Email       string `json:"email,omitempty"`
		Password    string `json:"password,omitempty"`
	}
	var reqBody RequestBody
	if err := c.BindJSON(&reqBody); err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to parse the request body: %v", err))
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	user := &models.User{
		Name:        reqBody.Name,
		DateOfBirth: reqBody.DateOfBirth,
		Address:     reqBody.Address,
		PrimaryRole: reqBody.PrimaryRole,
		Email:       reqBody.Email,
		Password:    reqBody.Password,
	}

	createdUser, err := h.authService.Register(c, user)
	if err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to register the user: %v", err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := utils.ResponseRenderer("Successfully registered the user", gin.H{
		"user": createdUser,
	})
	c.JSON(http.StatusOK, resp)
}
