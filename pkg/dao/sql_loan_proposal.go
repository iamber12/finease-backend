package dao

import (
	"context"
	"errors"
	"fmt"

	"bitbucket.com/finease/backend/pkg/db"
	"bitbucket.com/finease/backend/pkg/models"
	"gorm.io/gorm"
)

type sqlLoanProposal struct {
	sessionFactory db.SessionFactory
}

func NewSqlLoanProposalDao(factory db.SessionFactory) LoanProposal {
	return &sqlLoanProposal{sessionFactory: factory}
}

func (s *sqlLoanProposal) FindById(ctx context.Context, id string) (*models.LoanProposal, error) {
	tx := s.sessionFactory.New(ctx)
	var existingLoanProposal *models.LoanProposal
	err := tx.Where("uuid = ?", id).First(&existingLoanProposal).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find the loan proposal: %w", err)
		}
		return nil, fmt.Errorf("loan proposal not found")
	}
	return existingLoanProposal, nil
}

func (s *sqlLoanProposal) FindAll(ctx context.Context) ([]*models.LoanProposal, error) {
	tx := s.sessionFactory.New(ctx)
	var existingLoanProposals []*models.LoanProposal
	err := tx.Find(&existingLoanProposals).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find the loan proposals: %w", err)
		}
		return []*models.LoanProposal{}, nil
	}
	return existingLoanProposals, nil
}

func (s *sqlLoanProposal) FindByUserUuid(ctx context.Context, userUuid string) ([]*models.LoanProposal, error) {
	tx := s.sessionFactory.New(ctx)
	var existingLoanProposals []*models.LoanProposal
	err := tx.Where("user_uuid = ?", userUuid).Find(&existingLoanProposals).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find the loan proposals: %w", err)
		}
		return []*models.LoanProposal{}, nil
	}
	return existingLoanProposals, nil
}

func (s *sqlLoanProposal) Create(ctx context.Context, loanProposal *models.LoanProposal) (*models.LoanProposal, error) {
	tx := s.sessionFactory.New(ctx)

	if err := tx.Create(loanProposal).Error; err != nil {
		return nil, fmt.Errorf("unable to create the loan proposal in the DB: %w", err)
	}

	return s.FindById(ctx, loanProposal.Uuid)
}

func (s *sqlLoanProposal) Update(ctx context.Context, id string, patch *models.LoanProposal) (*models.LoanProposal, error) {
	tx := s.sessionFactory.New(ctx)

	var existingLoanProposal *models.LoanProposal
	err := tx.Where("uuid = ?", id).First(&existingLoanProposal).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find the existing loan proposal")
		}
		return nil, fmt.Errorf("loan proposal not found")
	}
	if err := tx.Model(existingLoanProposal).Where("uuid = ?", id).Updates(patch).Error; err != nil {
		return nil, fmt.Errorf("unable to update the loan proposal: %w", err)
	}
	return s.FindById(ctx, id)
}

func (s *sqlLoanProposal) Delete(ctx context.Context, id string) error {
	tx := s.sessionFactory.New(ctx)
	if err := tx.Unscoped().Where("uuid = ?", id).Delete(&models.LoanProposal{}).Error; err != nil {
		return fmt.Errorf("unable to delete the loan proposal: %w", err)
	}
	return nil
}
