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

type SupportTicket interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindByUserId(c *gin.Context)
	Delete(c *gin.Context)
}

var _ SupportTicket = supportTicketHandler{}

type supportTicketHandler struct {
	supportTicketService services.SupportTicket
}

func NewSupportTicketHandler(supportTicketService services.SupportTicket) SupportTicket {
	return &supportTicketHandler{supportTicketService: supportTicketService}
}

func (l supportTicketHandler) Create(c *gin.Context) {
	var reqBody api.SupportTicket

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

	inboundSupportTicketModel := api.MapSupportTicketApiToModel(&reqBody)
	inboundSupportTicketModel.UserUUID = userUuid

	createdSupportTicket, err := l.supportTicketService.Create(c, userUuid, inboundSupportTicketModel)
	if err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to create the support ticket: %v", err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	outboundSupportTicket := api.MapSupportTicketModelToApi(createdSupportTicket)

	resp := utils.ResponseRenderer("Support created successfully", gin.H{
		"support_ticket": outboundSupportTicket,
	})
	c.JSON(http.StatusOK, resp)
}

func (l supportTicketHandler) Delete(c *gin.Context) {
	supportTicketUuid := c.Param("support_ticket_uuid")
	user, ok := c.Keys["user"].(models.User)
	if !ok {
		resp := utils.ResponseRenderer("failed to parse the user uuid from the processed header")
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	userUuid := user.Uuid

	if err := l.supportTicketService.Delete(c, userUuid, supportTicketUuid); err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to delete the support ticket: %v", err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := utils.ResponseRenderer("Support ticket deleted successfully")
	c.JSON(http.StatusOK, resp)
}

func (l supportTicketHandler) Update(c *gin.Context) {
	supportTicketUuid := c.Param("support_ticket_uuid")
	user, ok := c.Keys["user"].(models.User)
	if !ok {
		resp := utils.ResponseRenderer("failed to parse the user uuid from the processed header")
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	userUuid := user.Uuid

	var reqBody api.SupportTicket
	if err := c.BindJSON(&reqBody); err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to parse the request body: %v", err))
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	inboundSupportTicketModel := api.MapSupportTicketApiToModel(&reqBody)
	updatedSupportTicket, err := l.supportTicketService.Update(c, userUuid, supportTicketUuid, inboundSupportTicketModel)
	if err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to update the support ticket: %v", err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	outboundSupportTicket := api.MapSupportTicketModelToApi(updatedSupportTicket)

	resp := utils.ResponseRenderer("Support ticket updated successfully", gin.H{
		"support_ticket": outboundSupportTicket,
	})
	c.JSON(http.StatusOK, resp)
}

func (l supportTicketHandler) FindById(c *gin.Context) {
	user, ok := c.Keys["user"].(models.User)
	if !ok {
		resp := utils.ResponseRenderer("failed to parse the user uuid from the processed header")
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	supportTicketUuid := c.Param("support_ticket_uuid")
	supportTicketResp, err := l.supportTicketService.FindById(c, supportTicketUuid)

	if err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to list your support ticket: %v", err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if supportTicketResp.UserUUID != user.Uuid {
		resp := utils.ResponseRenderer("failed to update the support ticket: not authorized to fetch the support ticket of some other user")
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	outboundSupportTicket := api.MapSupportTicketModelToApi(supportTicketResp)

	resp := utils.ResponseRenderer("Your support ticket fetched successfully", gin.H{
		"support_ticket": outboundSupportTicket,
	})
	c.JSON(http.StatusOK, resp)
}

func (l supportTicketHandler) FindByUserId(c *gin.Context) {
	user, ok := c.Keys["user"].(models.User)
	if !ok {
		resp := utils.ResponseRenderer("failed to parse the user uuid from the processed header")
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	userUuid := user.Uuid
	supportTickets, err := l.supportTicketService.FindByUserId(c, userUuid)

	if err != nil {
		resp := utils.ResponseRenderer(fmt.Sprintf("failed to list your support ticket: %v", err))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	outboundSupportTicket := make([]*api.SupportTicket, len(supportTickets))
	for i, supportTicket := range supportTickets {
		outboundSupportTicket[i] = api.MapSupportTicketModelToApi(supportTicket)
	}
	resp := utils.ResponseRenderer("Your support ticket fetched successfully", gin.H{
		"support_tickets": outboundSupportTicket,
	})
	c.JSON(http.StatusOK, resp)
}
