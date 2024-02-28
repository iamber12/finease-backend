package models

type LoanRequest struct {
	Generic
	UserUUID      string  `gorm:"not null"`
	Amount        float64 `gorm:"not null"`
	MaxInterest   float64 `gorm:"not null"`
	DurationToPay int64   `gorm:"not null"`
	Status        string  // enum: offered, available (potentially others as well)
	Description   string
}
