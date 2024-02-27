package api

import "bitbucket.com/finease/backend/pkg/models"

type UserList struct {
	Items []*User `json:"users,omitempty"`
}

type User struct {
	Uuid        string `json:"uuid,omitempty"`
	Name        string `json:"name,omitempty"`
	DateOfBirth string `json:"date_of_birth,omitempty"`
	Address     string `json:"address,omitempty"`
	PrimaryRole string `json:"primary_role,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
}

func MapUserRequestToModel(user *User) *models.User {
	return &models.User{
		Name:        user.Name,
		DateOfBirth: user.DateOfBirth,
		Address:     user.Address,
		PrimaryRole: user.PrimaryRole,
		Email:       user.Email,
		Password:    user.Password,
	}
}

func MapUserModelToRespone(user *models.User) *User {
	return &User{
		Uuid:        user.Uuid,
		Name:        user.Name,
		DateOfBirth: user.DateOfBirth,
		Address:     user.Address,
		PrimaryRole: user.PrimaryRole,
		Email:       user.Email,
		Password:    user.Password,
	}
}
