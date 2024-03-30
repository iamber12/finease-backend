package dao

import (
	"bitbucket.com/finease/backend/pkg/models"
	"context"
)

type LoanAgreement interface {
	Upsert(ctx context.Context, loanAgreement *models.LoanAgreement) (*models.LoanAgreement, error)
	FindOne(ctx context.Context, loanProposalUuid, loanRequestUuid string) (*models.LoanAgreement, error)
	Delete(ctx context.Context, loanProposalUuid, loanRequestUuid string) error
}
