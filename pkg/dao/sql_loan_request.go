package dao

import (
	"context"
	"errors"
	"fmt"

	"bitbucket.com/finease/backend/pkg/db"
	"bitbucket.com/finease/backend/pkg/models"
	"gorm.io/gorm"
)

type sqlLoanRequest struct {
	sessionFactory db.SessionFactory
}

func NewSqlLoanRequestDao(factory db.SessionFactory) LoanRequest {
	return &sqlLoanRequest{sessionFactory: factory}
}

func (s *sqlLoanRequest) FindById(ctx context.Context, id string) (*models.LoanRequest, error) {
	tx := s.sessionFactory.New(ctx)
	var existingLoanRequest *models.LoanRequest
	err := tx.Where("uuid = ?", id).First(&existingLoanRequest).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find the loan request: %w", err)
		}
		return nil, fmt.Errorf("loan request not found")
	}
	return existingLoanRequest, nil
}

func (s *sqlLoanRequest) FindByUserId(ctx context.Context, userUuid string) ([]*models.LoanRequest, error) {
	tx := s.sessionFactory.New(ctx)
	var existingLoanRequest []*models.LoanRequest
	err := tx.Where("user_uuid = ?", userUuid).Find(&existingLoanRequest).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find the loan request: %w", err)
		}
		return nil, fmt.Errorf("loan request not found")
	}
	return existingLoanRequest, nil
}

func (s *sqlLoanRequest) FindAll(ctx context.Context) ([]*models.LoanRequest, error) {
	tx := s.sessionFactory.New(ctx)
	var existingLoanRequests []*models.LoanRequest
	err := tx.Find(&existingLoanRequests).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find the loan requests: %w", err)
		}
		return []*models.LoanRequest{}, nil
	}
	return existingLoanRequests, nil
}

func (s *sqlLoanRequest) Create(ctx context.Context, loanRequest *models.LoanRequest) (*models.LoanRequest, error) {
	tx := s.sessionFactory.New(ctx)

	if err := tx.Create(loanRequest).Error; err != nil {
		return nil, fmt.Errorf("unable to create the loan request in the DB: %w", err)
	}

	return s.FindById(ctx, loanRequest.Uuid)
}

func (s *sqlLoanRequest) Update(ctx context.Context, id string, patch *models.LoanRequest) (*models.LoanRequest, error) {
	tx := s.sessionFactory.New(ctx)

	var existingLoanRequest *models.LoanRequest
	err := tx.Where("uuid = ?", id).First(&existingLoanRequest).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find the existing loan request")
		}
		return nil, fmt.Errorf("loan request not found")
	}
	if err := tx.Model(existingLoanRequest).Where("uuid = ?", id).Updates(patch).Error; err != nil {
		return nil, fmt.Errorf("unable to update the loan request: %w", err)
	}
	return s.FindById(ctx, id)
}

func (s *sqlLoanRequest) Delete(ctx context.Context, id string) error {
	tx := s.sessionFactory.New(ctx)
	if err := tx.Unscoped().Where("uuid = ?", id).Delete(&models.LoanRequest{}).Error; err != nil {
		return fmt.Errorf("unable to delete the loan request: %w", err)
	}
	return nil
}
