package services

import (
	"context"
	// "time"

	"bitbucket.com/finease/backend/pkg/models"
	// "github.com/google/uuid"
)

type LoanRequest interface {
	Create(ctx context.Context, loanRequest *models.LoanRequest) (*models.LoanRequest, error)
	Update(ctx context.Context, loanRequest *models.LoanRequest) (*models.LoanRequest, error)
	Delete(ctx context.Context, id string) (string, error)
	FindById(ctx context.Context, id string) (*models.LoanRequest, error)
	FindAll(ctx context.Context) ([]*models.LoanRequest, error)
}

type loanRequestService struct {
	loanRequestDao models.LoanRequest
}

// func (l loanRequestService) Create(ctx context.Context, loanRequest *models.LoanRequest) (*models.LoanRequest, error) {
// 	var err error

// 	loanRequest.CreatedAt, loanRequest.UpdatedAt = time.Now(), time.Now()
// 	loanRequest.Uuid = uuid.New().String()
// }
