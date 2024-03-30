package models

import (
	"gorm.io/datatypes"
)

type LoanProposal struct {
	Generic
	UserUUID             string
	AmountStart          float64
	AmountEnd            float64
	MinInterest          float64
	MaxInterest          float64
	MaxReturnDuration    int64
	MinReturnDuration    int64
	LenderAccountDetails datatypes.JSONMap
	Status               string
	Description          string
}

type LoanProposalStatus string

const (
	LOAN_PROPOSAL_AVAILABLE LoanProposalStatus = "available"
	LOAN_PROPOSAL_GRANTED   LoanProposalStatus = "granted"
	LOAN_PROPOSAL_OFFERED   LoanProposalStatus = "offered"
)
