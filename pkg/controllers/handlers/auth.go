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

	jwtToken, err := h.authService.Login(c, reqBody.Email, reqBody.Password)
	if err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to authenticate the user %v", err))
		c.JSON(http.StatusForbidden, resp)
		return
	}

	resp := utils.ResponseRenderer("Validation successful", gin.H{
		"jwtToken": jwtToken,
	})
	c.JSON(http.StatusOK, resp)
}

func (h authHandler) Register(c *gin.Context) {
	var reqBody api.User
	if err := c.BindJSON(&reqBody); err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to parse the request body: %v", err))
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	inboundUserModel := api.MapUserApiToModel(&reqBody)

	createdUser, err := h.authService.Register(c, inboundUserModel)
	if err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to register the user: %v", err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	outboundUserResponse := api.MapUserModelToApi(createdUser)

	resp := utils.ResponseRenderer("Successfully registered the user", gin.H{
		"user": outboundUserResponse,
	})
	c.JSON(http.StatusOK, resp)
}
