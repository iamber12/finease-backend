package models

type LoanRequest struct {
	Generic
	UserUUID      string  `gorm:"not null"`
	Amount        float64 `gorm:"not null"`
	MinInterest   float64 `gorm:"not null"`
	MaxInterest   float64 `gorm:"not null"`
	DurationToPay int64   `gorm:"not null"`
	Status        *string // enum: offered, available (potentially others as well)
	ProposalUuid  *string
	Description   string
}

type LoanRequestStatus string

const (
	LOAN_REQUEST_ACCEPTED LoanRequestStatus = "accepted"
	LOAN_REQUEST_GRANTED  LoanRequestStatus = "granted"
	LOAN_REQUEST_REJECTED LoanRequestStatus = "rejected"
)
