package migrations

// Migrations should NEVER use types from other packages. Types can change
// and then migrations run on a _new_ database will fail or behave unexpectedly.
// Instead of importing types, always re-create the type in the migration, as
// is done here, even though the same type is defined in pkg/api

import (
	"gorm.io/gorm"

	"github.com/go-gormigrate/gormigrate/v2"
)

func removeLoanProposalsUniqueUserUuidIndex() *gormigrate.Migration {
	type LoanProposal struct {
	}
	return &gormigrate.Migration{
		ID: "202402212106",
		Migrate: func(tx *gorm.DB) error {
			if tx.Migrator().HasIndex(&LoanProposal{}, "loan_proposal_user_uuid_unique_idx") {
				return tx.Migrator().DropIndex(&LoanProposal{}, "loan_proposal_user_uuid_unique_idx")
			}
			return nil
		},
	}
}
