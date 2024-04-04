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

func addFinancialTransactionsTable() *gormigrate.Migration {
	type FinancialTransactions struct {
		Model
		LoanProposalUuid string  `gorm:"not null"`
		LoanRequestUuid  string  `gorm:"not null"`
		BorrowerUuid     string  `gorm:"not null"` // for caching purposes
		LenderUuid       string  `gorm:"not null"` // for caching purposes
		Amount           float64 `gorm:"not null"`
		Payer            string  `gorm:"not null"` // enum: borrower/lender
		DateOffered      time.Time
	}
	return &gormigrate.Migration{
		ID: "202404022127",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().AutoMigrate(&FinancialTransactions{})
		},
	}
}
