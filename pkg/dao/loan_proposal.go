package dao

import (
	"context"

	"bitbucket.com/finease/backend/pkg/models"
)

type LoanProposal interface {
	Create(ctx context.Context, loanProposal *models.LoanProposal) (*models.LoanProposal, error)
	FindById(ctx context.Context, id string) (*models.LoanProposal, error)
	FindAll(ctx context.Context) ([]*models.LoanProposal, error)
	FindByUserUuid(ctx context.Context, userUuid string) ([]*models.LoanProposal, error)
	Update(ctx context.Context, id string, patch *models.LoanProposal) (*models.LoanProposal, error)
	Delete(ctx context.Context, id string) error
}
