package models

import "time"

type FinancialTransaction struct {
	Generic
	LoanProposalUuid string    `gorm:"not null"`
	LoanRequestUuid  string    `gorm:"not null"`
	BorrowerUuid     string    `gorm:"not null"` // for caching purposes
	LenderUuid       string    `gorm:"not null"` // for caching purposes
	PayerType        PayerType `gorm:"not null"`
	Amount           float64   `gorm:"not null"`
	DateOffered      time.Time
}

type PayerType string

const (
	PAYER_TYPE_LENDER   PayerType = "lender"
	PAYER_TYPE_BORROWER PayerType = "borrower"
)
