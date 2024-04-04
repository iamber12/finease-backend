package api

import (
	"bitbucket.com/finease/backend/pkg/models"
	"time"
)

type FinancialTransaction struct {
	Uuid             string    `json:"uuid,omitempty"`
	LoanProposalUuid string    `json:"loan_proposal_uuid,omitempty"`
	LoanRequestUuid  string    `json:"loan_request_uuid,omitempty"`
	BorrowerUuid     string    `json:"borrower_uuid,omitempty"`
	LenderUuid       string    `json:"lender_uuid,omitempty"`
	PayerType        string    `json:"payer_type,omitempty"`
	Amount           float64   `json:"amount,omitempty"`
	DateOffered      time.Time `json:"date_offered,omitempty"`
}

func MapFinancialTransactionModelToApi(transaction *models.FinancialTransaction) *FinancialTransaction {
	return &FinancialTransaction{
		Uuid:             transaction.Uuid,
		LoanProposalUuid: transaction.LoanProposalUuid,
		LoanRequestUuid:  transaction.LoanRequestUuid,
		BorrowerUuid:     transaction.BorrowerUuid,
		LenderUuid:       transaction.LenderUuid,
		PayerType:        string(transaction.PayerType),
		Amount:           transaction.Amount,
		DateOffered:      transaction.DateOffered,
	}
}

func MapFinancialTransactionApiToModel(transaction *FinancialTransaction) *models.FinancialTransaction {
	return &models.FinancialTransaction{
		Generic: models.Generic{
			Uuid: transaction.Uuid,
		},
		LoanProposalUuid: transaction.LoanProposalUuid,
		LoanRequestUuid:  transaction.LoanRequestUuid,
		BorrowerUuid:     transaction.BorrowerUuid,
		LenderUuid:       transaction.LenderUuid,
		PayerType:        models.PayerType(transaction.PayerType),
		Amount:           transaction.Amount,
		DateOffered:      transaction.DateOffered,
	}
}
