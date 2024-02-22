package migrations

// Migrations should NEVER use types from other packages. Types can change
// and then migrations run on a _new_ database will fail or behave unexpectedly.
// Instead of importing types, always re-create the type in the migration, as
// is done here, even though the same type is defined in pkg/api

import (
	"gorm.io/gorm"

	"github.com/go-gormigrate/gormigrate/v2"
)

func addLoadProposalsTable() *gormigrate.Migration {
	type LoanProposal struct {
		Model
		UserUUID string `gorm:"uniqueIndex:loan_proposal_user_uuid_unique_idx;not null"`

		AmountStart float64 `gorm:"not null"`
		AmountEnd   float64 `gorm:"not null"`

		MinInterest float64 `gorm:"not null"`
		MaxInterest float64 `gorm:"not null"`

		MaxReturnDuration int64 `gorm:"not null"`
		MinReturnDuration int64 `gorm:"not null"`

		LenderAccountDetails map[string]any `gorm:"type:jsonb"`

		Status string // enum: offered, available (potentially others as well)

		Description string
	}
	return &gormigrate.Migration{
		ID: "202402212017",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().AutoMigrate(&LoanProposal{})
		},
	}
}
