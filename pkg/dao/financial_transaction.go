package dao

import (
	"bitbucket.com/finease/backend/pkg/models"
	"context"
)

type FinancialTransaction interface {
	Create(ctx context.Context, transaction *models.FinancialTransaction) (*models.FinancialTransaction, error)
	FindByLenderUuid(ctx context.Context, lenderUuid string) ([]*models.FinancialTransaction, error)
	FindByBorrowerUuid(ctx context.Context, borrowerUuid string) ([]*models.FinancialTransaction, error)
	FindByLoanProposalUuid(ctx context.Context, userUuid, proposalUuid string) ([]*models.FinancialTransaction, error)
	FindByLoanRequestUuid(ctx context.Context, userUuid, requestUuid string) ([]*models.FinancialTransaction, error)

	FindByUserUuid(ctx context.Context, userUuid string) ([]*models.FinancialTransaction, error)
	FindReceived(ctx context.Context, userUuid string) ([]*models.FinancialTransaction, error)
	FindSent(ctx context.Context, userUuid string) ([]*models.FinancialTransaction, error)
	FindByLoanAgreement(ctx context.Context, agreement *models.LoanAgreement) (*models.FinancialTransaction, error)
	Update(ctx context.Context, id string, patch *models.FinancialTransaction) (*models.FinancialTransaction, error)
	Delete(ctx context.Context, transactionUuid string) error
}
