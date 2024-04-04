package dao

import (
	"bitbucket.com/finease/backend/pkg/db"
	"bitbucket.com/finease/backend/pkg/models"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type sqlLoanAgreement struct {
	sessionFactory db.SessionFactory
}

func NewSqlLoanAgreementDao(factory db.SessionFactory) LoanAgreement {
	return &sqlLoanAgreement{sessionFactory: factory}
}

func (s *sqlLoanAgreement) Upsert(ctx context.Context, agreement *models.LoanAgreement) (*models.LoanAgreement, error) {
	tx := s.sessionFactory.New(ctx)

	if err := tx.Create(agreement).Error; err != nil {
		return nil, fmt.Errorf("unable to create the loan agreement in the DB: %w", err)
	}

	err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "loan_proposal_uuid"}, {Name: "loan_request_uuid"},
		}, // if ON CONFLICT on the primary key loan_proposal_uuid, loan_request_uuid, then....
		DoNothing: true, // ... skip
	}).Create(agreement).Error // ... else create
	if err != nil {
		return nil, fmt.Errorf("unable to create the loan agreement in the DB: %w", err)
	}

	return s.FindOne(ctx, agreement.LoanProposalUuid, agreement.LoanRequestUuid)
}

func (s *sqlLoanAgreement) FindOne(ctx context.Context, loanProposalUuid, loanRequestUuid string) (*models.LoanAgreement, error) {
	tx := s.sessionFactory.New(ctx)

	var agreement models.LoanAgreement

	if err := tx.Where("loan_proposal_uuid = ? AND loan_request_uuid = ?", loanProposalUuid, loanRequestUuid).Find(&agreement).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find the loan requests: %w", err)
		}
		return nil, fmt.Errorf("loan agreement not found")
	}
	return &agreement, nil
}

func (s *sqlLoanAgreement) Delete(ctx context.Context, loanProposalUuid, loanRequestUuid string) error {
	tx := s.sessionFactory.New(ctx)

	if err := tx.Where("loan_proposal_uuid = ? AND loan_request_uuid = ?", loanProposalUuid, loanRequestUuid).
		Unscoped().Delete(&models.LoanAgreement{}).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to delete loan agreement: %w", err)
	}

	return nil
}
