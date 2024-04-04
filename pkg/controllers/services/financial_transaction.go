package services

import (
	"bitbucket.com/finease/backend/pkg/dao"
	"bitbucket.com/finease/backend/pkg/models"
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type FinancialTransaction interface {
	Create(ctx context.Context, transaction *models.FinancialTransaction) (*models.FinancialTransaction, error)
	FindAll(ctx context.Context, userUuid string) ([]*models.FinancialTransaction, error)
	FindReceived(ctx context.Context, userUuid string) ([]*models.FinancialTransaction, error)
	FindSent(ctx context.Context, userUuid string) ([]*models.FinancialTransaction, error)
	FindByLoanProposalUuid(ctx context.Context, proposerUuid string, proposalUuid string) ([]*models.FinancialTransaction, error)
	FindByLoanRequestUuid(ctx context.Context, requesterUuid string, requestUuid string) ([]*models.FinancialTransaction, error)
}

type financialTransactionService struct {
	financialTransactionsDao dao.FinancialTransaction
}

func NewFinancialTransactionService(financialTransactionsDao dao.FinancialTransaction) FinancialTransaction {
	return &financialTransactionService{
		financialTransactionsDao: financialTransactionsDao,
	}
}

func (f financialTransactionService) Create(ctx context.Context, transaction *models.FinancialTransaction) (*models.FinancialTransaction, error) {
	transaction.CreatedAt, transaction.UpdatedAt = time.Now(), time.Now()
	transaction.Uuid = uuid.New().String()
	createdFinancialTransaction, err := f.financialTransactionsDao.Create(ctx, transaction)
	if err != nil {
		return nil, fmt.Errorf("failed to create the financial transaction: %w", err)
	}
	return createdFinancialTransaction, nil
}

func (f financialTransactionService) FindAll(ctx context.Context, userUuid string) ([]*models.FinancialTransaction, error) {
	transactions, err := f.financialTransactionsDao.FindByUserUuid(ctx, userUuid)
	if err != nil {
		return []*models.FinancialTransaction{}, fmt.Errorf("failed to fetch your financial transactions: %w", err)
	}
	return transactions, nil
}

func (f financialTransactionService) FindReceived(ctx context.Context, userUuid string) ([]*models.FinancialTransaction, error) {
	transactions, err := f.financialTransactionsDao.FindReceived(ctx, userUuid)
	if err != nil {
		return []*models.FinancialTransaction{}, fmt.Errorf("failed to fetch the financial transactions you received: %w", err)
	}
	return transactions, nil
}

func (f financialTransactionService) FindSent(ctx context.Context, userUuid string) ([]*models.FinancialTransaction, error) {
	transactions, err := f.financialTransactionsDao.FindSent(ctx, userUuid)
	if err != nil {
		return []*models.FinancialTransaction{}, fmt.Errorf("failed to fetch the financial transactions you received: %w", err)
	}
	return transactions, nil
}

func (f financialTransactionService) FindByLoanProposalUuid(ctx context.Context, proposerUuid string, proposalUuid string) ([]*models.FinancialTransaction, error) {
	transactions, err := f.financialTransactionsDao.FindByLoanProposalUuid(ctx, proposerUuid, proposalUuid)
	if err != nil {
		return []*models.FinancialTransaction{}, fmt.Errorf("failed to fetch the financial transactions by this loan proposal: %w", err)
	}
	return transactions, nil
}

func (f financialTransactionService) FindByLoanRequestUuid(ctx context.Context, requesterUuid string, requestUuid string) ([]*models.FinancialTransaction, error) {
	transactions, err := f.financialTransactionsDao.FindByLoanRequestUuid(ctx, requesterUuid, requestUuid)
	if err != nil {
		return []*models.FinancialTransaction{}, fmt.Errorf("failed to fetch the financial transactions by this loan request: %w", err)
	}
	return transactions, nil
}
