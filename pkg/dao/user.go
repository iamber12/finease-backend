package dao

import (
	"context"

	"bitbucket.com/finease/backend/pkg/models"
)

type User interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	FindById(ctx context.Context, id string) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	Update(ctx context.Context, id string, patch *models.User) (*models.User, error)
	Delete(ctx context.Context, id string) error
}
