package services

import (
	"fmt"

	"bitbucket.com/finease/backend/pkg/dao"
	"bitbucket.com/finease/backend/pkg/models"
	"github.com/gin-gonic/gin"
)

type User interface {
	FindById(c *gin.Context, userUuid string) (*models.User, error)
	Update(c *gin.Context, userUuid string, patch *models.User) (*models.User, error)
}

type userService struct {
	userDao dao.User
}

func NewUserService(userDao dao.User) User {
	return &userService{userDao: userDao}
}

func (u userService) FindById(ctx *gin.Context, userUuid string) (*models.User, error) {
	user, err := u.userDao.FindById(ctx, userUuid)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}

	return user, nil
}

func (u userService) Update(ctx *gin.Context, userUuid string, patch *models.User) (*models.User, error) {
	user, err := u.userDao.Update(ctx, userUuid, patch)
	if err != nil {
		return nil, fmt.Errorf("failed to update user details: %w", err)
	}

	return user, nil
}
