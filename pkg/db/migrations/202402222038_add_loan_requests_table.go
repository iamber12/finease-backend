package migrations

// Migrations should NEVER use types from other packages. Types can change
// and then migrations run on a _new_ database will fail or behave unexpectedly.
// Instead of importing types, always re-create the type in the migration, as
// is done here, even though the same type is defined in pkg/api

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func addLoadRequestTable() *gormigrate.Migration {
	type LoanRequest struct {
		Model
		UserUUID      string  `gorm:"not null"`
		Amount        float64 `gorm:"not null"`
		MinInterest   float64 `gorm:"not null"`
		MaxInterest   float64 `gorm:"not null"`
		DurationToPay int64   `gorm:"not null"`
		Status        string  // enum: offered, available (potentially others as well)
		Description   string
	}
	return &gormigrate.Migration{
		ID: "202402222038",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().AutoMigrate(&LoanRequest{})
		},
	}
}
