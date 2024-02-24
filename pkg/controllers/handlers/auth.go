package handlers

import (
	"bitbucket.com/finease/backend/pkg/controllers/api"
	"bitbucket.com/finease/backend/pkg/controllers/services"
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
	var reqBody api.User
	if err := c.BindJSON(&reqBody); err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to parse the request body: %v", err))
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	validators := []Validator{
		ValidateEmail,
	}

	for _, validator := range validators {
		if !validator(reqBody) {
			resp := utils.ResponseRenderer("invalid request body found")
			c.JSON(http.StatusBadRequest, resp)
			return
		}
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
