package services

import (
	"context"
	"fmt"
	"time"

	// "time"

	"bitbucket.com/finease/backend/pkg/dao"
	"bitbucket.com/finease/backend/pkg/models"
	"github.com/google/uuid"
	// "github.com/google/uuid"
)

type LoanRequest interface {
	Create(ctx context.Context, loanRequest *models.LoanRequest) (*models.LoanRequest, error)
	Update(ctx context.Context, ownerUserUuid, id string, patch *models.LoanRequest) (*models.LoanRequest, error)
	Delete(ctx context.Context, ownerUserUuid, id string) error
	FindByUserId(ctx context.Context, ownerUserUuid string) ([]*models.LoanRequest, error)
	FindById(ctx context.Context, id string) (*models.LoanRequest, error)
}

type loanRequestService struct {
	loanRequestDao dao.LoanRequest
}

func (l loanRequestService) Create(ctx context.Context, loanRequest *models.LoanRequest) (*models.LoanRequest, error) {
	loanRequest.CreatedAt, loanRequest.UpdatedAt = time.Now(), time.Now()
	loanRequest.Uuid = uuid.New().String()

	createdLoanRequest, err := l.loanRequestDao.Create(ctx, loanRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to create the loan request: %w", err)
	}

	return createdLoanRequest, err
}

func (l loanRequestService) Update(ctx context.Context, ownerUserUuid, loanRequestUuid string, patch *models.LoanRequest) (*models.LoanRequest, error) {
	if patch.Uuid != "" {
		return nil, fmt.Errorf("not allowed to update UUID")
	}

	loanRequestFound, err := l.loanRequestDao.FindById(ctx, loanRequestUuid)
	if err != nil {
		return nil, fmt.Errorf("failed to find the loan request corresponding to the provided loan request uuid: %w", err)
	}

	if ownerUserUuid != loanRequestFound.UserUUID {
		return nil, fmt.Errorf("not authorized to update the loan request of some other user")
	}

	updatedLoanRequest, err := l.loanRequestDao.Update(ctx, loanRequestUuid, patch)
	if err != nil {
		return nil, fmt.Errorf("failed to update the loan request: %w", err)
	}
	return updatedLoanRequest, nil
}

func (l loanRequestService) Delete(ctx context.Context, ownerUserUuid, loanRequestUuid string) error {
	loanRequestFound, err := l.loanRequestDao.FindById(ctx, loanRequestUuid)
	if err != nil {
		return fmt.Errorf("failed to find the loan request corresponding to the provided loan request uuid: %w", err)
	}

	if ownerUserUuid != loanRequestFound.UserUUID {
		return fmt.Errorf("not authorized to delete the loan request of some other user")
	}

	if err := l.loanRequestDao.Delete(ctx, loanRequestUuid); err != nil {
		return fmt.Errorf("failed to delete the loan request: %w", err)
	}
	return nil
}

func (l loanRequestService) FindById(ctx context.Context, id string) (*models.LoanRequest, error) {
	loanRequests, err := l.loanRequestDao.FindById(ctx, id)
	if err != nil {
		return &models.LoanRequest{}, fmt.Errorf("failed to fetch your loan requests: %w", err)
	}
	return loanRequests, nil
}

func (l loanRequestService) FindByUserId(ctx context.Context, ownerUserUuid string) ([]*models.LoanRequest, error) {
	loanRequests, err := l.loanRequestDao.FindByUserId(ctx, ownerUserUuid)
	if err != nil {
		return []*models.LoanRequest{}, fmt.Errorf("failed to fetch your loan requests: %w", err)
	}
	return loanRequests, nil
}
