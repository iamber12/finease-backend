package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Uuid        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	Name        string
	DateOfBirth string
	Address     string
	PrimaryRole string
	Email       string
	Password    string
	Active      *bool
}
