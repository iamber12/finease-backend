package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var MigrationList = []*gormigrate.Migration{
	addUserTable(),
	addLoadProposalsTable(),
	removeLoanProposalsUniqueUserUuidIndex(),
	addLoanRequestTable(),
	addLoanAgreementsTable(),
	addFinancialTranscationsTable(),
	usersTableEmailUniqueIndex(),
	addProposalUuidToLoanRequestsTable(),
	userActiveField(),
	removeOldLoanAgreementsTable(),
	addNewLoanAgreementsTable(),
	addSupportTicketTable(),
	addFinancialTransactionsTable(),
}

// Model represents the base model struct. All entities will have this struct embedded.
type Model struct {
	Uuid      string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
