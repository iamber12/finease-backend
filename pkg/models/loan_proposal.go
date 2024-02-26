package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type LoanProposal struct {
	Uuid                 string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt
	UserUUID             string
	AmountStart          float64
	AmountEnd            float64
	MinInterest          float64
	MaxInterest          float64
	MaxReturnDuration    int64
	MinReturnDuration    int64
	LenderAccountDetails datatypes.JSONMap
	Status               string // enum: offered, available (potentially others as well)
	Description          string
}
