package migrations

// Migrations should NEVER use types from other packages. Types can change
// and then migrations run on a _new_ database will fail or behave unexpectedly.
// Instead of importing types, always re-create the type in the migration, as
// is done here, even though the same type is defined in pkg/api

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func addLoadAgreementsTable() *gormigrate.Migration {
	type LoanAgreement struct {
		Model
		LoanAgreementUUID    string         `gorm:"not null"`
		Amount               float64        `gorm:"not null"`
		LenderAccountDetails map[string]any `gorm:"type:jsonb"`
		Date                 time.Time
	}
	return &gormigrate.Migration{
		ID: "202402222150",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().AutoMigrate(&LoanAgreement{})
		},
	}
}
