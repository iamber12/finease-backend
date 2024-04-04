package services

import (
	"context"
	"fmt"
	"time"

	"bitbucket.com/finease/backend/pkg/dao"
	"bitbucket.com/finease/backend/pkg/models"
	"github.com/google/uuid"
)

type SupportTicket interface {
	Create(ctx context.Context, userUuid string, supportTicket *models.SupportTicket) (*models.SupportTicket, error)
	Update(ctx context.Context, userUuid, id string, patch *models.SupportTicket) (*models.SupportTicket, error)
	Delete(ctx context.Context, userUuid, id string) error
	FindByUserId(ctx context.Context, userUuid string) ([]*models.SupportTicket, error)
	FindById(ctx context.Context, id string) (*models.SupportTicket, error)
}

type supportTicketService struct {
	supportTicketDao dao.SupportTicket
	userDao          dao.User
}

func NewSupportTicketService(supportTicketDao dao.SupportTicket, userDao dao.User) SupportTicket {
	return &supportTicketService{
		supportTicketDao: supportTicketDao,
		userDao:          userDao,
	}
}

func (l supportTicketService) Create(ctx context.Context, userUuid string, supportTicket *models.SupportTicket) (*models.SupportTicket, error) {
	supportTicket.CreatedAt, supportTicket.UpdatedAt = time.Now(), time.Now()
	supportTicket.Uuid = uuid.New().String()

	createdSupportTicket, err := l.supportTicketDao.Create(ctx, supportTicket)
	if err != nil {
		return nil, fmt.Errorf("failed to create the support ticket: %w", err)
	}

	return createdSupportTicket, err
}

func (l supportTicketService) Update(ctx context.Context, userUuid, supportTicketUuid string, patch *models.SupportTicket) (*models.SupportTicket, error) {
	if patch.Uuid != "" {
		return nil, fmt.Errorf("not allowed to update UUID")
	}

	supportTicketFound, err := l.supportTicketDao.FindById(ctx, supportTicketUuid)
	if err != nil {
		return nil, fmt.Errorf("failed to find the support ticket corresponding to the provided support ticket uuid: %w", err)
	}

	if userUuid != supportTicketFound.UserUUID {
		return nil, fmt.Errorf("not authorized to update the support ticket of some other user")
	}

	updatedSupportTicket, err := l.supportTicketDao.Update(ctx, supportTicketUuid, patch)
	if err != nil {
		return nil, fmt.Errorf("failed to update the support ticket: %w", err)
	}
	return updatedSupportTicket, nil
}

func (l supportTicketService) Delete(ctx context.Context, userUuid, supportTicketUuid string) error {
	supportTicketFound, err := l.supportTicketDao.FindById(ctx, supportTicketUuid)
	if err != nil {
		return fmt.Errorf("failed to find the support ticket corresponding to the provided support ticket uuid: %w", err)
	}

	if userUuid != supportTicketFound.UserUUID {
		return fmt.Errorf("not authorized to delete the support ticket of some other user")
	}

	if err := l.supportTicketDao.Delete(ctx, supportTicketUuid); err != nil {
		return fmt.Errorf("failed to delete the support ticket: %w", err)
	}
	return nil
}

func (l supportTicketService) FindById(ctx context.Context, id string) (*models.SupportTicket, error) {
	supportTickets, err := l.supportTicketDao.FindById(ctx, id)
	if err != nil {
		return &models.SupportTicket{}, fmt.Errorf("failed to fetch your support tickets: %w", err)
	}
	return supportTickets, nil
}

func (l supportTicketService) FindByUserId(ctx context.Context, userUuid string) ([]*models.SupportTicket, error) {
	supportTickets, err := l.supportTicketDao.FindByUserId(ctx, userUuid)
	if err != nil {
		return []*models.SupportTicket{}, fmt.Errorf("failed to fetch your support tickets: %w", err)
	}
	return supportTickets, nil
}
