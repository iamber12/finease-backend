package api

import (
	"bitbucket.com/finease/backend/pkg/models"
)

type SupportTicket struct {
	Uuid        string `json:"uuid,omitempty"`
	UserUUID    string `json:"user_uuid,omitempty"`
	Status      string `json:"status,omitempty"`
	Description string `json:"description,omitempty"`
	Subject     string `json:"subject,omitempty"`
}

func MapSupportTicketModelToApi(supportTicket *models.SupportTicket) *SupportTicket {
	return &SupportTicket{
		Uuid:        supportTicket.Uuid,
		UserUUID:    supportTicket.UserUUID,
		Status:      supportTicket.Status,
		Description: supportTicket.Description,
		Subject:     supportTicket.Subject,
	}
}

func MapSupportTicketApiToModel(supportTicket *SupportTicket) *models.SupportTicket {
	return &models.SupportTicket{
		UserUUID:    supportTicket.UserUUID,
		Status:      supportTicket.Status,
		Description: supportTicket.Description,
		Subject:     supportTicket.Subject,
	}
}
