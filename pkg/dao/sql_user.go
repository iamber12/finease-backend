package dao

import (
	"bitbucket.com/finease/backend/pkg/db"
	"bitbucket.com/finease/backend/pkg/models"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type sqlUser struct {
	sessionFactory db.SessionFactory
}

func NewSqlUserDao(factory db.SessionFactory) User {
	return &sqlUser{sessionFactory: factory}
}

func (s *sqlUser) Create(ctx context.Context, user *models.User) (*models.User, error) {
	tx := s.sessionFactory.New(ctx)

	var existingUser *models.User
	err := tx.Where("uuid = ?", user.Uuid).First(&existingUser).Error
	shouldCreate := false
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to check the error")
		}
		shouldCreate = true
	}

	if !shouldCreate {
		return existingUser, nil
	}

	if err := tx.Create(user).Error; err != nil {
		return nil, fmt.Errorf("unable to create the user in the DB: %w", err)
	}

	return user, nil
}

func (s *sqlUser) FindById(ctx context.Context, id string) (*models.User, error) {
	tx := s.sessionFactory.New(ctx)
	var existingUser *models.User
	err := tx.Where("uuid = ?", id).First(&existingUser).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to check the error")
		}
		return nil, nil
	}
	return existingUser, nil

}

func (s *sqlUser) Update(ctx context.Context, id string, clauses map[string]interface{}) (*models.User, error) {
	tx := s.sessionFactory.New(ctx)

	var existingUser *models.User
	err := tx.Where("uuid = ?", id).First(&existingUser).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to check the error")
		}
		return nil, fmt.Errorf("user not found")
	}
	return existingUser, nil
}

func (s *sqlUser) Delete(ctx context.Context, id string) error {
	tx := s.sessionFactory.New(ctx)
	if err := tx.Unscoped().Where("uuid = ?", id).Delete(&models.User{}).Error; err != nil {
		return fmt.Errorf("unable to delete the user: %w", err)
	}
	return nil
}
