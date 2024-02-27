package services

import (
	"context"
	"fmt"
	"time"

	"bitbucket.com/finease/backend/pkg/dao"
	"bitbucket.com/finease/backend/pkg/models"
	"github.com/google/uuid"
)

type LoanProposal interface {
	Create(ctx context.Context, loanProposal *models.LoanProposal) (*models.LoanProposal, error)
	Update(ctx context.Context, ownerUserUuid, id string, patch *models.LoanProposal) (*models.LoanProposal, error)
	Delete(ctx context.Context, ownerUserUuid, loanProposalUuid string) error
	Find(ctx context.Context, ownerUserUuid string) ([]*models.LoanProposal, error)
}

type loanProposalService struct {
	userDao         dao.User
	loanProposalDao dao.LoanProposal
}

func NewLoanProposalService(loanProposalDao dao.LoanProposal, userDao dao.User) LoanProposal {
	return &loanProposalService{loanProposalDao: loanProposalDao, userDao: userDao}
}

func (l loanProposalService) Create(ctx context.Context, loanProposal *models.LoanProposal) (*models.LoanProposal, error) {
	loanProposal.CreatedAt, loanProposal.UpdatedAt = time.Now(), time.Now()
	loanProposal.Uuid = uuid.New().String()

	createdLoanProposal, err := l.loanProposalDao.Create(ctx, loanProposal)
	if err != nil {
		return nil, fmt.Errorf("failed to create the loan proposal: %w", err)
	}
	return createdLoanProposal, nil
}

func (l loanProposalService) Update(ctx context.Context, ownerUserUuid, loanProposalUuid string, patch *models.LoanProposal) (*models.LoanProposal, error) {
	if patch.Uuid != "" {
		return nil, fmt.Errorf("not allowed to update UUID")
	}

	loanProposalFound, err := l.loanProposalDao.FindById(ctx, loanProposalUuid)
	if err != nil {
		return nil, fmt.Errorf("failed to find the loan proposal corresponding to the provided loan proposal uuid: %w", err)
	}

	if ownerUserUuid != loanProposalFound.UserUUID {
		return nil, fmt.Errorf("not authorized to update the loan proposal of some other user")
	}

	updatedLoanProposal, err := l.loanProposalDao.Update(ctx, loanProposalUuid, patch)
	if err != nil {
		return nil, fmt.Errorf("failed to update the loan proposal: %w", err)
	}
	return updatedLoanProposal, nil
}

func (l loanProposalService) Delete(ctx context.Context, ownerUserUuid, loanProposalUuid string) error {
	loanProposalFound, err := l.loanProposalDao.FindById(ctx, loanProposalUuid)
	if err != nil {
		return fmt.Errorf("failed to find the loan proposal corresponding to the provided loan proposal uuid: %w", err)
	}

	if ownerUserUuid != loanProposalFound.UserUUID {
		return fmt.Errorf("not authorized to delete the loan proposal of some other user")
	}

	if err := l.loanProposalDao.Delete(ctx, loanProposalUuid); err != nil {
		return fmt.Errorf("failed to delete the loan proposal: %w", err)
	}
	return nil
}

func (l loanProposalService) Find(ctx context.Context, ownerUserUuid string) ([]*models.LoanProposal, error) {
	proposals, err := l.loanProposalDao.FindByUserUuid(ctx, ownerUserUuid)
	if err != nil {
		return []*models.LoanProposal{}, fmt.Errorf("failed to fetch your loan proposals: %w", err)
	}
	return proposals, nil
}
