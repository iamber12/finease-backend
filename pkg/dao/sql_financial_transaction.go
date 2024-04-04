package dao

import (
	"bitbucket.com/finease/backend/pkg/db"
	"bitbucket.com/finease/backend/pkg/models"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type sqlFinancialTransaction struct {
	sessionFactory db.SessionFactory
}

func NewSqlFinancialTransactionDao(factory db.SessionFactory) FinancialTransaction {
	return &sqlFinancialTransaction{sessionFactory: factory}
}

func (s *sqlFinancialTransaction) Create(ctx context.Context, transaction *models.FinancialTransaction) (*models.FinancialTransaction, error) {
	tx := s.sessionFactory.New(ctx)

	if err := tx.Create(transaction).Error; err != nil {
		return nil, fmt.Errorf("unable to create the financial transaction in the DB: %w", err)
	}

	return s.findByUuid(ctx, transaction.Uuid)
}

func (s *sqlFinancialTransaction) FindByLenderUuid(ctx context.Context, lenderUuid string) ([]*models.FinancialTransaction, error) {
	tx := s.sessionFactory.New(ctx)
	var transactions []*models.FinancialTransaction
	err := tx.Where("lender_uuid = ?", lenderUuid).Find(&transactions).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find the financial transaction: %w", err)
		}
		return []*models.FinancialTransaction{}, nil
	}
	return transactions, nil
}

func (s *sqlFinancialTransaction) FindByBorrowerUuid(ctx context.Context, borrowerUuid string) ([]*models.FinancialTransaction, error) {
	tx := s.sessionFactory.New(ctx)
	var transactions []*models.FinancialTransaction
	err := tx.Where("borrower_uuid = ?", borrowerUuid).Find(&transactions).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find the financial transactions: %w", err)
		}
		return []*models.FinancialTransaction{}, nil
	}
	return transactions, nil
}

func (s *sqlFinancialTransaction) FindByUserUuid(ctx context.Context, userUuid string) ([]*models.FinancialTransaction, error) {
	tx := s.sessionFactory.New(ctx)
	var transactions []*models.FinancialTransaction
	err := tx.Where("borrower_uuid = ? OR lender_uuid = ?", userUuid, userUuid).Find(&transactions).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find the financial transactions: %w", err)
		}
		return []*models.FinancialTransaction{}, nil
	}
	return transactions, nil
}

func (s *sqlFinancialTransaction) FindReceived(ctx context.Context, userUuid string) ([]*models.FinancialTransaction, error) {
	tx := s.sessionFactory.New(ctx)
	var transactions []*models.FinancialTransaction
	err := tx.Where("(borrower_uuid = ? AND payer_type = ?) OR (lender_uuid = ? AND payer_type = ?)",
		userUuid, models.PAYER_TYPE_LENDER, userUuid, models.PAYER_TYPE_BORROWER).Find(&transactions).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find the financial transactions: %w", err)
		}
		return []*models.FinancialTransaction{}, nil
	}
	return transactions, nil
}

func (s *sqlFinancialTransaction) FindSent(ctx context.Context, userUuid string) ([]*models.FinancialTransaction, error) {
	tx := s.sessionFactory.New(ctx)
	var transactions []*models.FinancialTransaction
	err := tx.Where("(borrower_uuid = ? AND payer_type = ?) OR (lender_uuid = ? AND payer_type = ?)",
		userUuid, models.PAYER_TYPE_BORROWER, userUuid, models.PAYER_TYPE_LENDER).Find(&transactions).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find the financial transactions: %w", err)
		}
		return []*models.FinancialTransaction{}, nil
	}
	return transactions, nil
}

func (s *sqlFinancialTransaction) FindByLoanProposalUuid(ctx context.Context, userUuid, proposalUuid string) ([]*models.FinancialTransaction, error) {
	tx := s.sessionFactory.New(ctx)
	var transactions []*models.FinancialTransaction
	err := tx.Where("loan_proposal_uuid = ? AND lender_uuid = ?", proposalUuid, userUuid).Find(&transactions).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find the financial transactions: %w", err)
		}
		return []*models.FinancialTransaction{}, nil
	}
	return transactions, nil
}

func (s *sqlFinancialTransaction) FindByLoanRequestUuid(ctx context.Context, userUuid, requestUuid string) ([]*models.FinancialTransaction, error) {
	tx := s.sessionFactory.New(ctx)
	var transactions []*models.FinancialTransaction
	err := tx.Where("loan_request_uuid = ? AND borrower_uuid = ?", requestUuid, userUuid).Find(&transactions).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find the financial transactions: %w", err)
		}
		return []*models.FinancialTransaction{}, nil
	}
	return transactions, nil
}

func (s *sqlFinancialTransaction) FindByLoanAgreement(ctx context.Context, agreement *models.LoanAgreement) (*models.FinancialTransaction, error) {
	tx := s.sessionFactory.New(ctx)
	var transaction models.FinancialTransaction
	err := tx.Where("loan_proposal_uuid = ? AND loan_request_uuid = ?", agreement.LoanProposalUuid, agreement.LoanRequestUuid).First(&transaction).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find the financial transaction: %w", err)
		}
		return nil, fmt.Errorf("financial transaction not found")
	}
	return &transaction, nil
}

func (s *sqlFinancialTransaction) findByUuid(ctx context.Context, transactionUuid string) (*models.FinancialTransaction, error) {
	tx := s.sessionFactory.New(ctx)
	var transaction models.FinancialTransaction
	err := tx.Where("uuid = ?", transactionUuid).First(&transaction).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find the financial transaction: %w", err)
		}
		return nil, fmt.Errorf("financial transaction not found")
	}
	return &transaction, nil
}

func (s *sqlFinancialTransaction) Update(ctx context.Context, id string, patch *models.FinancialTransaction) (*models.FinancialTransaction, error) {
	tx := s.sessionFactory.New(ctx)

	var existingTransaction models.FinancialTransaction
	err := tx.Where("uuid = ?", id).First(&existingTransaction).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find the existing financial transaction")
		}
		return nil, fmt.Errorf("financial transaction not found")
	}
	if err := tx.Model(&existingTransaction).Where("uuid = ?", id).Updates(patch).Error; err != nil {
		return nil, fmt.Errorf("unable to update the financial transaction: %w", err)
	}
	return s.findByUuid(ctx, id)
}

func (s *sqlFinancialTransaction) Delete(ctx context.Context, transactionUuid string) error {
	tx := s.sessionFactory.New(ctx)
	if err := tx.Unscoped().Where("uuid = ?", transactionUuid).Delete(&models.FinancialTransaction{}).Error; err != nil {
		return fmt.Errorf("unable to delete the financial transaction: %w", err)
	}
	return nil
}
