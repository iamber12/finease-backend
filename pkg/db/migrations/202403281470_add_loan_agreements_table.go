package migrations

// Migrations should NEVER use types from other packages. Types can change
// and then migrations run on a _new_ database will fail or behave unexpectedly.
// Instead of importing types, always re-create the type in the migration, as
// is done here, even though the same type is defined in pkg/api

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func addNewLoanAgreementsTable() *gormigrate.Migration {
	type LoanAgreement struct {
		LoanProposalUuid string `gorm:"primaryKey"`
		LoanRequestUuid  string `gorm:"primaryKey"`

		LenderUuid   string `gorm:"not null"`
		BorrowerUuid string `gorm:"not null"`

		Amount   float64 `gorm:"not null"`
		Interest float64 `gorm:"not null"`
		Duration int64   `gorm:"not null"`
	}
	return &gormigrate.Migration{
		ID: "202403281470",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().AutoMigrate(&LoanAgreement{})
		},
	}
}
