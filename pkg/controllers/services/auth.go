package services

import (
	"bitbucket.com/finease/backend/pkg/dao"
	"bitbucket.com/finease/backend/pkg/models"
	"bitbucket.com/finease/backend/pkg/utils"
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Auth interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
}

type authService struct {
	userDao dao.User
}

func NewAuthService(userDao dao.User) Auth {
	return authService{userDao: userDao}
}

func (a authService) Register(ctx context.Context, user *models.User) (*models.User, error) {
	var err error

	user.CreatedAt, user.UpdatedAt = time.Now(), time.Now()
	user.Uuid = uuid.New().String()
	user.Password, err = utils.Hash(user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash the incoming password: %w", err)
	}

	createdUser, err := a.userDao.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create the user: %w", err)
	}
	return createdUser, nil
}
