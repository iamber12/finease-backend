package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func addSupportTicketTable() *gormigrate.Migration {
	type SupportTicket struct {
		Model
		UserUUID    string `gorm:"not null"`
		Subject     string `gorm:"not null"`
		Description string `gorm:"not null"`
		Status      string `gorm:"not null"`
	}
	return &gormigrate.Migration{
		ID: "202403302328",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().AutoMigrate(&SupportTicket{})
		},
	}
}
