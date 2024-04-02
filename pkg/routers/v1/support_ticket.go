package v1

import (
	"bitbucket.com/finease/backend/pkg/controllers/handlers"
	"bitbucket.com/finease/backend/pkg/controllers/services"
	"github.com/gin-gonic/gin"
)

func SetupSupportTicketRouter(parentRouter *gin.RouterGroup, supportTicket services.SupportTicket, additionalMiddlewares ...gin.HandlerFunc) {
	supportTicketRouter := parentRouter.Group("/support_ticket")
	supportTicketHandler := handlers.NewSupportTicketHandler(supportTicket)

	supportTicketRouter.Use(additionalMiddlewares...)

	supportTicketRouter.POST("/", supportTicketHandler.Create)
	supportTicketRouter.GET("/:support_ticket_uuid", supportTicketHandler.FindById)
	supportTicketRouter.GET("/my", supportTicketHandler.FindByUserId)
	supportTicketRouter.PUT("/:support_ticket_uuid", supportTicketHandler.Update)
	supportTicketRouter.DELETE("/:support_ticket_uuid", supportTicketHandler.Delete)
}
