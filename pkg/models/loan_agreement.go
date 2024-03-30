package models

type LoanAgreement struct {
	LoanProposalUuid string
	LoanRequestUuid  string
	LenderUuid       string
	BorrowerUuid     string
	Amount           float64
	Interest         float64
	Duration         int64
}
