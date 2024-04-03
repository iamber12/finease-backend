package handlers

import (
	"fmt"
	"net/http"

	"bitbucket.com/finease/backend/pkg/controllers/api"
	"bitbucket.com/finease/backend/pkg/controllers/services"
	"bitbucket.com/finease/backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

type Auth interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type authHandler struct {
	authService services.Auth
}

func NewAuthHandler(authService services.Auth) Auth {
	return &authHandler{authService: authService}
}

func (h authHandler) Login(c *gin.Context) {
	var reqBody api.LoginRequestBody
	if err := c.BindJSON(&reqBody); err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to parse the request body: %v", err))
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	jwtToken, user, err := h.authService.Login(c, reqBody.Email, reqBody.Password)
	if err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to authenticate the user %v", err))
		c.JSON(http.StatusForbidden, resp)
		return
	}

	resp := utils.ResponseRenderer("Validation successful", gin.H{
		"jwt_token": jwtToken,
		"user":      api.MapUserModelToRespone(user),
	})
	c.JSON(http.StatusOK, resp)
}

func (h authHandler) Register(c *gin.Context) {
	var reqBody api.User
	isActive := true

	if err := c.BindJSON(&reqBody); err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to parse the request body: %v", err))
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	validators := []Validator{
		ValidateEmail,
		ValidateDOB,
	}

	for _, validator := range validators {
		if !validator(reqBody) {
			resp := utils.ResponseRenderer("invalid request body found")
			c.JSON(http.StatusBadRequest, resp)
			return
		}
	}

	inboundUserModel := api.MapUserRequestToModel(&reqBody)
	inboundUserModel.Active = &isActive

	createdUser, err := h.authService.Register(c, *inboundUserModel)
	if err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to register the user: %v", err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	outboundUserResponse := api.MapUserModelToRespone(createdUser)

	resp := utils.ResponseRenderer("Successfully registered the user", gin.H{
		"user": outboundUserResponse,
	})
	c.JSON(http.StatusOK, resp)
}
