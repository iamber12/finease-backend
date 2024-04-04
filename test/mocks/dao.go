package mocks

import (
	"context"
	"fmt"

	"bitbucket.com/finease/backend/pkg/dao"
	"bitbucket.com/finease/backend/pkg/models"
)

type sqlUserDaoMock struct {
	db map[string]*models.User
}

func NewUserDaoMock() dao.User {
	return &sqlUserDaoMock{db: map[string]*models.User{}}
}

func (s sqlUserDaoMock) Create(ctx context.Context, user *models.User) (*models.User, error) {
	existingUser, _ := s.FindById(ctx, user.Uuid)
	if existingUser != nil {
		return nil, fmt.Errorf("existing user with the same uuid found")
	}
	existingUser, _ = s.FindByEmail(ctx, user.Email)
	if existingUser != nil {
		return nil, fmt.Errorf("existing user with the same email found")
	}
	s.db[user.Uuid] = user
	return user, nil
}

func (s sqlUserDaoMock) FindById(ctx context.Context, id string) (*models.User, error) {
	userFound, ok := s.db[id]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	return userFound, nil
}

func (s sqlUserDaoMock) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	for _, user := range s.db {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func (s sqlUserDaoMock) Update(ctx context.Context, id string, patch *models.User) (*models.User, error) {
	existingUser, err := s.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	if patch.Uuid != "" {
		existingUser.Uuid = patch.Uuid
	}
	if patch.Name != "" {
		existingUser.Name = patch.Name
	}
	if patch.DateOfBirth != "" {
		existingUser.DateOfBirth = patch.DateOfBirth
	}
	if patch.Address != "" {
		existingUser.Address = patch.Address
	}
	if patch.PrimaryRole != "" {
		existingUser.PrimaryRole = patch.PrimaryRole
	}
	if patch.Email != "" {
		existingUser.Email = patch.Email
	}
	if patch.Password != "" {
		existingUser.Password = patch.Password
	}

	s.db[id] = existingUser
	return existingUser, nil
}

func (s sqlUserDaoMock) Delete(ctx context.Context, id string) error {
	delete(s.db, id)
	return nil
}
