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
	var existingLoanProposal *models.LoanRequest
	err := tx.Where("uuid = ?", id).First(&existingLoanProposal).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find the loan proposal: %w", err)
		}
		return nil, fmt.Errorf("loan proposal not found")
	}
	return existingLoanProposal, nil
}

func (s *sqlLoanRequest) FindAll(ctx context.Context) ([]*models.LoanRequest, error) {
	tx := s.sessionFactory.New(ctx)
	var existingLoanProposals []*models.LoanRequest
	err := tx.Find(&existingLoanProposals).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find the loan proposals: %w", err)
		}
		return []*models.LoanRequest{}, nil
	}
	return existingLoanProposals, nil
}

func (s *sqlLoanRequest) Create(ctx context.Context, loanProposal *models.LoanRequest) (*models.LoanRequest, error) {
	tx := s.sessionFactory.New(ctx)

	if err := tx.Create(loanProposal).Error; err != nil {
		return nil, fmt.Errorf("unable to create the loan proposal in the DB: %w", err)
	}

	return s.FindById(ctx, loanProposal.Uuid)
}

func (s *sqlLoanRequest) Update(ctx context.Context, id string, patch *models.LoanRequest) (*models.LoanRequest, error) {
	tx := s.sessionFactory.New(ctx)

	var existingLoanProposal *models.LoanRequest
	err := tx.Where("uuid = ?", id).First(&existingLoanProposal).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find the existing loan proposal")
		}
		return nil, fmt.Errorf("loan proposal not found")
	}
	if err := tx.Model(existingLoanProposal).Updates(patch).Error; err != nil {
		return nil, fmt.Errorf("unable to update the loan proposal: %w", err)
	}
	return s.FindById(ctx, id)
}

func (s *sqlLoanRequest) Delete(ctx context.Context, id string) error {
	tx := s.sessionFactory.New(ctx)
	if err := tx.Unscoped().Where("uuid = ?", id).Delete(&models.LoanRequest{}).Error; err != nil {
		return fmt.Errorf("unable to delete the loan proposal: %w", err)
	}
	return nil
}
