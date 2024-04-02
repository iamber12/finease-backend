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

type User interface {
	Update(c *gin.Context)
	FindById(c *gin.Context)
	Deactivate(c *gin.Context)
}

type userHandler struct {
	userService services.User
}

func NewUserHandler(userService services.User) User {
	return &userHandler{userService: userService}
}

func (u userHandler) FindById(c *gin.Context) {
	user, ok := c.Keys["user"].(models.User)
	if !ok {
		resp := utils.ResponseRenderer("failed to parse the user uuid from the processed header")
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	userUuid := c.Param("user_uuid")
	if userUuid == "" {
		userUuid = user.Uuid
	}

	userResp, err := u.userService.FindById(c, userUuid)
	if err != nil {
		resp := utils.ResponseRenderer(err.Error())
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	outboundUser := api.MapUserModelToRespone(userResp)
	resp := utils.ResponseRenderer("User succesfully fetched", gin.H{
		"user": outboundUser,
	})

	c.JSON(http.StatusOK, resp)
}

func (u userHandler) callUpdateService(c *gin.Context, userUuid string, inboundUserModel *models.User, isUpdate bool) {
	updatedUser, err := u.userService.Update(c, userUuid, inboundUserModel)
	if err != nil {
		resp := utils.ResponseRenderer(err.Error())
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	outboundUser := api.MapUserModelToRespone(updatedUser)

	var resp gin.H
	if isUpdate {
		resp = utils.ResponseRenderer("User succesfully updated", gin.H{
			"user": outboundUser,
		})
	} else {
		resp = utils.ResponseRenderer("User deactivated")
	}

	c.JSON(http.StatusOK, resp)
}

func (u userHandler) Update(c *gin.Context) {
	var reqBody api.User
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

	validators := make([]Validator, 0)

	if reqBody.Email != "" {
		validators = append(validators, ValidateEmail)
	}
	if reqBody.DateOfBirth != "" {
		validators = append(validators, ValidateDOB)
	}

	for _, validator := range validators {
		if !validator(reqBody) {
			resp := utils.ResponseRenderer("invalid request body found")
			c.JSON(http.StatusBadRequest, resp)
			return
		}
	}

	inboundUserModel := api.MapUserRequestToModel(&reqBody)
	u.callUpdateService(c, user.Uuid, inboundUserModel, true)
}

func (u userHandler) Deactivate(c *gin.Context) {
	isActive := false

	user, ok := c.Keys["user"].(models.User)

	if !ok {
		resp := utils.ResponseRenderer("failed to parse the user uuid from the processed header")
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	userToUpdate := models.User{Active: &isActive}

	u.callUpdateService(c, user.Uuid, &userToUpdate, false)
}
