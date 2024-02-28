package api

import "bitbucket.com/finease/backend/pkg/models"

type LoanRequest struct {
	Uuid          string  `json:"uuid,omitempty"`
	UserUUID      string  `json:"user_uuid,omitempty"`
	Amount        float64 `json:"amount,omitempty"`
	MaxInterest   float64 `json:"max_interest,omitempty"`
	DurationToPay int64   `json:"duration,omitempty"`
	Status        string  `json:"status,omitempty"`
	Description   string  `json:"description,omitempty"`
}

func MapLoanRequestModelToApi(loanRequest *models.LoanRequest) *LoanRequest {
	return &LoanRequest{
		Uuid:          loanRequest.Uuid,
		UserUUID:      loanRequest.UserUUID,
		Amount:        loanRequest.Amount,
		MaxInterest:   loanRequest.MaxInterest,
		DurationToPay: loanRequest.DurationToPay,
		Status:        loanRequest.Status,
		Description:   loanRequest.Description,
	}
}

func MapLoanRequestApiToModel(loanRequest *LoanRequest) *LoanRequest {
	return &LoanRequest{
		UserUUID:      loanRequest.UserUUID,
		Amount:        loanRequest.Amount,
		MaxInterest:   loanRequest.MaxInterest,
		DurationToPay: loanRequest.DurationToPay,
		Status:        loanRequest.Status,
		Description:   loanRequest.Description,
	}
}
