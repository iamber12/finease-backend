package models

import (
	"time"

	"gorm.io/gorm"
)

type Generic struct {
	Uuid      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
