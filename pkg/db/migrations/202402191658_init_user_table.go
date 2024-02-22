package migrations

// Migrations should NEVER use types from other packages. Types can change
// and then migrations run on a _new_ database will fail or behave unexpectedly.
// Instead of importing types, always re-create the type in the migration, as
// is done here, even though the same type is defined in pkg/api

import (
	"gorm.io/gorm"

	"github.com/go-gormigrate/gormigrate/v2"
)

func addUserTable() *gormigrate.Migration {
	type User struct {
		Model
		Name        string
		DateOfBirth string
		Address     string
		PrimaryRole string // enum: borrower/lender
		Email       string
		Password    string
	}
	return &gormigrate.Migration{
		ID: "202402191658",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().AutoMigrate(&User{})
		},
	}
}
