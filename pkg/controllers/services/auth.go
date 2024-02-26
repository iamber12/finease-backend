package services

import (
	"context"
	"fmt"
	"time"

	"bitbucket.com/finease/backend/pkg/dao"
	"bitbucket.com/finease/backend/pkg/models"
	"bitbucket.com/finease/backend/pkg/utils"
	"github.com/google/uuid"
)

type Auth interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
	Login(ctx context.Context, email string, password string) (string, *models.User, error)
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

func (a authService) Login(ctx context.Context, email string, password string) (string, *models.User, error) {
	var err error
	userDetails, err := a.userDao.FindByEmail(ctx, email)

	if err != nil {
		return "", nil, fmt.Errorf("failed to get the user: %w", err)
	}

	if !utils.ValidatePassword(password, userDetails.Password) {
		return "", nil, fmt.Errorf("invalid password")
	}

	token, err := utils.GenerateJWT(userDetails.Uuid, "success")
	if err != nil {
		return "", nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return token, userDetails, nil
}
