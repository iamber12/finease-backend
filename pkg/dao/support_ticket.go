package dao

import (
	"context"

	"bitbucket.com/finease/backend/pkg/models"
)

type SupportTicket interface {
	Create(ctx context.Context, supportTicket *models.SupportTicket) (*models.SupportTicket, error)
	Update(ctx context.Context, id string, patch *models.SupportTicket) (*models.SupportTicket, error)
	Delete(ctx context.Context, id string) error
	FindByUserId(ctx context.Context, userUuid string) ([]*models.SupportTicket, error)
	FindById(ctx context.Context, id string) (*models.SupportTicket, error)
}
