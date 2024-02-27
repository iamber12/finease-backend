package api

import (
	"bitbucket.com/finease/backend/pkg/models"
)

type LoanProposal struct {
	Uuid                 string         `json:"uuid,omitempty"`
	UserUUID             string         `json:"user_uuid,omitempty"`
	AmountStart          float64        `json:"amount_start,omitempty"`
	AmountEnd            float64        `json:"amount_end,omitempty"`
	MinInterest          float64        `json:"min_interest,omitempty"`
	MaxInterest          float64        `json:"max_interest,omitempty"`
	MaxReturnDuration    int64          `json:"max_return_duration,omitempty"`
	MinReturnDuration    int64          `json:"min_return_duration,omitempty"`
	LenderAccountDetails map[string]any `json:"lender_account_details,omitempty"`
	Status               string         `json:"status,omitempty"` // enum: offered, available (potentially others as well)
	Description          string         `json:"description,omitempty"`
}

func MapLoanProposalRequestToModel(loanProposal *LoanProposal) *models.LoanProposal {
	return &models.LoanProposal{
		AmountStart:          loanProposal.AmountStart,
		AmountEnd:            loanProposal.AmountEnd,
		MinInterest:          loanProposal.MinInterest,
		MaxInterest:          loanProposal.MaxInterest,
		MaxReturnDuration:    loanProposal.MaxReturnDuration,
		MinReturnDuration:    loanProposal.MinReturnDuration,
		LenderAccountDetails: loanProposal.LenderAccountDetails,
		Status:               loanProposal.Status,
		Description:          loanProposal.Description,
	}
}

func MapLoanProposalModelToResponse(loanProposal *models.LoanProposal) *LoanProposal {
	return &LoanProposal{
		Uuid:                 loanProposal.Uuid,
		UserUUID:             loanProposal.UserUUID,
		AmountStart:          loanProposal.AmountStart,
		AmountEnd:            loanProposal.AmountEnd,
		MinInterest:          loanProposal.MinInterest,
		MaxInterest:          loanProposal.MaxInterest,
		MaxReturnDuration:    loanProposal.MaxReturnDuration,
		MinReturnDuration:    loanProposal.MinReturnDuration,
		LenderAccountDetails: loanProposal.LenderAccountDetails,
		Status:               loanProposal.Status,
		Description:          loanProposal.Description,
	}
}
