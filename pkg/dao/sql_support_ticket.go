package dao

import (
	"context"
	"errors"
	"fmt"

	"bitbucket.com/finease/backend/pkg/db"
	"bitbucket.com/finease/backend/pkg/models"
	"gorm.io/gorm"
)

type sqlSupportTicket struct {
	sessionFactory db.SessionFactory
}

func NewSqlSupportTicketDao(factory db.SessionFactory) SupportTicket {
	return &sqlSupportTicket{sessionFactory: factory}
}

func (s *sqlSupportTicket) FindById(ctx context.Context, id string) (*models.SupportTicket, error) {
	tx := s.sessionFactory.New(ctx)
	var existingSupportTicket *models.SupportTicket
	err := tx.Where("uuid = ?", id).First(&existingSupportTicket).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find the support ticket: %w", err)
		}
		return nil, fmt.Errorf("support ticket not found")
	}
	return existingSupportTicket, nil
}

func (s *sqlSupportTicket) FindByUserId(ctx context.Context, userUuid string) ([]*models.SupportTicket, error) {
	tx := s.sessionFactory.New(ctx)
	var existingSupportTicket []*models.SupportTicket
	err := tx.Where("user_uuid = ?", userUuid).Find(&existingSupportTicket).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find the support ticket: %w", err)
		}
		return nil, fmt.Errorf("support ticket not found")
	}
	return existingSupportTicket, nil
}

func (s *sqlSupportTicket) Create(ctx context.Context, supportTicket *models.SupportTicket) (*models.SupportTicket, error) {
	tx := s.sessionFactory.New(ctx)

	if err := tx.Create(supportTicket).Error; err != nil {
		return nil, fmt.Errorf("unable to create the support ticket in the DB: %w", err)
	}

	return s.FindById(ctx, supportTicket.Uuid)
}

func (s *sqlSupportTicket) Update(ctx context.Context, id string, patch *models.SupportTicket) (*models.SupportTicket, error) {
	tx := s.sessionFactory.New(ctx)

	var existingSupportTicket *models.SupportTicket
	err := tx.Where("uuid = ?", id).First(&existingSupportTicket).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find the existing support ticket")
		}
		return nil, fmt.Errorf("support ticket not found")
	}
	if err := tx.Model(existingSupportTicket).Where("uuid = ?", id).Updates(patch).Error; err != nil {
		return nil, fmt.Errorf("unable to update the support ticket: %w", err)
	}
	return s.FindById(ctx, id)
}

func (s *sqlSupportTicket) Delete(ctx context.Context, id string) error {
	tx := s.sessionFactory.New(ctx)
	if err := tx.Unscoped().Where("uuid = ?", id).Delete(&models.SupportTicket{}).Error; err != nil {
		return fmt.Errorf("unable to delete the support ticket: %w", err)
	}
	return nil
}
