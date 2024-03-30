package services

import (
	"bitbucket.com/finease/backend/pkg/utils"
	"context"
	"fmt"
	"time"

	"bitbucket.com/finease/backend/pkg/dao"
	"bitbucket.com/finease/backend/pkg/models"
	"github.com/google/uuid"
)

type LoanProposal interface {
	Create(ctx context.Context, loanProposal *models.LoanProposal) (*models.LoanProposal, error)
	Update(ctx context.Context, ownerUserUuid, id string, patch *models.LoanProposal) (*models.LoanProposal, error)
	Delete(ctx context.Context, ownerUserUuid, loanProposalUuid string) error
	FindAvailable(ctx context.Context) ([]*models.LoanProposal, error)
	FindMine(ctx context.Context, ownerUserUuid string) ([]*models.LoanProposal, error)

	OfferGrant(ctx context.Context, granterUserUuid, proposalUuid, requestUuid string) error
	RevokeGrant(ctx context.Context, revokerUserUuid, proposalUuid, requestUuid string) error
	AcceptGrant(ctx context.Context, acceptorUserUuid, proposalUuid, requestUuid string) error
	RejectGrant(ctx context.Context, rejectorUserUuid, proposalUuid, requestUuid string) error
}

type loanProposalService struct {
	userDao          dao.User
	loanProposalDao  dao.LoanProposal
	loanRequestDao   dao.LoanRequest
	loanAgreementDao dao.LoanAgreement
}

func NewLoanProposalService(loanProposalDao dao.LoanProposal,
	loanRequestDao dao.LoanRequest,
	loanAgreementDao dao.LoanAgreement,
	userDao dao.User) LoanProposal {
	return &loanProposalService{
		loanProposalDao:  loanProposalDao,
		loanRequestDao:   loanRequestDao,
		loanAgreementDao: loanAgreementDao,
		userDao:          userDao,
	}
}

func (l loanProposalService) Create(ctx context.Context, loanProposal *models.LoanProposal) (*models.LoanProposal, error) {
	loanProposal.CreatedAt, loanProposal.UpdatedAt = time.Now(), time.Now()
	loanProposal.Uuid = uuid.New().String()

	createdLoanProposal, err := l.loanProposalDao.Create(ctx, loanProposal)
	if err != nil {
		return nil, fmt.Errorf("failed to create the loan proposal: %w", err)
	}
	return createdLoanProposal, nil
}

func (l loanProposalService) Update(ctx context.Context, ownerUserUuid, loanProposalUuid string, patch *models.LoanProposal) (*models.LoanProposal, error) {
	if patch.Uuid != "" {
		return nil, fmt.Errorf("not allowed to update UUID")
	}

	loanProposalFound, err := l.loanProposalDao.FindById(ctx, loanProposalUuid)
	if err != nil {
		return nil, fmt.Errorf("failed to find the loan proposal corresponding to the provided loan proposal uuid: %w", err)
	}

	if ownerUserUuid != loanProposalFound.UserUUID {
		return nil, fmt.Errorf("not authorized to update the loan proposal of some other user")
	}

	updatedLoanProposal, err := l.loanProposalDao.Update(ctx, loanProposalUuid, patch)
	if err != nil {
		return nil, fmt.Errorf("failed to update the loan proposal: %w", err)
	}
	return updatedLoanProposal, nil
}

func (l loanProposalService) Delete(ctx context.Context, ownerUserUuid, loanProposalUuid string) error {
	loanProposalFound, err := l.loanProposalDao.FindById(ctx, loanProposalUuid)
	if err != nil {
		return fmt.Errorf("failed to find the loan proposal corresponding to the provided loan proposal uuid: %w", err)
	}

	if ownerUserUuid != loanProposalFound.UserUUID {
		return fmt.Errorf("not authorized to delete the loan proposal of some other user")
	}

	if err := l.loanProposalDao.Delete(ctx, loanProposalUuid); err != nil {
		return fmt.Errorf("failed to delete the loan proposal: %w", err)
	}
	return nil
}

func (l loanProposalService) FindAvailable(ctx context.Context) ([]*models.LoanProposal, error) {
	proposals, err := l.loanProposalDao.FindAll(ctx)
	if err != nil {
		return []*models.LoanProposal{}, fmt.Errorf("failed to fetch available loan proposals: %w", err)
	}
	return proposals, nil
}

func (l loanProposalService) FindMine(ctx context.Context, ownerUserUuid string) ([]*models.LoanProposal, error) {
	proposals, err := l.loanProposalDao.FindByUserUuid(ctx, ownerUserUuid)
	if err != nil {
		return []*models.LoanProposal{}, fmt.Errorf("failed to fetch available loan proposals: %w", err)
	}
	return proposals, nil
}

func (l loanProposalService) OfferGrant(ctx context.Context, granterUserUuid, proposalUuid, requestUuid string) error {
	loanProposal, err := l.loanProposalDao.FindById(ctx, proposalUuid)
	if err != nil {
		return fmt.Errorf("failed to get the loan proposal: %w", err)
	}

	if loanProposal.UserUUID != granterUserUuid {
		return fmt.Errorf("cannot propose someone else's proposals")
	}

	loanRequestsAgainstProposal, err := l.loanRequestDao.FindByProposalId(ctx, proposalUuid)
	if err != nil {
		return fmt.Errorf("failed to check the potential loan requests associated with this loan proposals: %w", err)
	}

	for _, loanRequest := range loanRequestsAgainstProposal {
		if utils.FromPtr(loanRequest.Status) == string(models.LOAN_REQUEST_ACCEPTED) {
			return fmt.Errorf("cannot grant a proposal which is already approved against a loan request")
		}
		if utils.FromPtr(loanRequest.Status) == string(models.LOAN_REQUEST_GRANTED) {
			if loanRequest.Uuid == requestUuid {
				return nil
			}
			return fmt.Errorf("cannot grant a proposal to multiple loan requests")
		}
	}

	loanRequest, err := l.loanRequestDao.FindById(ctx, requestUuid)
	if err != nil {
		return fmt.Errorf("failed to get the loan request: %w", err)
	}

	if utils.FromPtr(loanRequest.ProposalUuid) != "" {
		return fmt.Errorf("cannot grant loan originally requested against a specific proposal")
	}

	if _, err := l.loanRequestDao.Update(ctx, requestUuid, &models.LoanRequest{
		ProposalUuid: &proposalUuid,
		Status:       utils.ToPtr(string(models.LOAN_REQUEST_GRANTED)),
	}); err != nil {
		return fmt.Errorf("failed to update the status of the loan request: %w", err)
	}

	if _, err := l.loanProposalDao.Update(ctx, proposalUuid, &models.LoanProposal{
		Status: string(models.LOAN_PROPOSAL_GRANTED),
	}); err != nil {
		return fmt.Errorf("failed to update the status of the loan proposal: %w", err)
	}

	return nil
}

func (l loanProposalService) RevokeGrant(ctx context.Context, revokerUserUuid, proposalUuid, requestUuid string) error {
	loanProposal, err := l.loanProposalDao.FindById(ctx, proposalUuid)
	if err != nil {
		return fmt.Errorf("failed to get the loan proposal: %w", err)
	}

	if loanProposal.UserUUID != revokerUserUuid {
		return fmt.Errorf("cannot revoke on behalf of someone else's proposals")
	}

	loanRequest, err := l.loanRequestDao.FindById(ctx, requestUuid)
	if err != nil {
		return fmt.Errorf("failed to get the loan request: %w", err)
	}

	if utils.FromPtr(loanRequest.ProposalUuid) != proposalUuid {
		return fmt.Errorf("cannot revoke a grant not referencing to this proposal uuid")
	}
	if utils.FromPtr(loanRequest.Status) != string(models.LOAN_REQUEST_GRANTED) {
		return fmt.Errorf("loan request not found to be granted in the first place")
	}

	if _, err := l.loanRequestDao.Update(ctx, requestUuid, &models.LoanRequest{
		ProposalUuid: utils.ToPtr(""),
		Status:       utils.ToPtr(""),
	}); err != nil {
		return fmt.Errorf("failed to update the status of the loan request: %w", err)
	}

	if _, err := l.loanProposalDao.Update(ctx, proposalUuid, &models.LoanProposal{
		Status: string(models.LOAN_PROPOSAL_AVAILABLE),
	}); err != nil {
		return fmt.Errorf("failed to update the status of the loan proposal: %w", err)
	}

	return nil
}

func (l loanProposalService) AcceptGrant(ctx context.Context, acceptorUserUuid, proposalUuid, requestUuid string) error {
	loanRequest, err := l.loanRequestDao.FindById(ctx, requestUuid)
	if err != nil {
		return fmt.Errorf("failed to get the loan request: %w", err)
	}
	loanProposal, err := l.loanProposalDao.FindById(ctx, proposalUuid)
	if err != nil {
		return fmt.Errorf("failed to get the loan proposal: %w", err)
	}

	if loanRequest.UserUUID != acceptorUserUuid {
		return fmt.Errorf("cannot accept someone else's grant")
	}
	if utils.FromPtr(loanRequest.Status) != string(models.LOAN_REQUEST_GRANTED) {
		return fmt.Errorf("request not found to have a grant associated with it")
	}
	if utils.FromPtr(loanRequest.ProposalUuid) != proposalUuid {
		return fmt.Errorf("request grant not found to be associated with the provided proposal uuid")
	}

	_, err = l.loanRequestDao.Update(ctx, requestUuid, &models.LoanRequest{
		Status: utils.ToPtr(string(models.LOAN_REQUEST_ACCEPTED)),
	})
	if err != nil {
		return fmt.Errorf("failed to update the status of the loan request: %w", err)
	}

	if _, err := l.loanProposalDao.Update(ctx, proposalUuid, &models.LoanProposal{
		Status: string(models.LOAN_PROPOSAL_OFFERED),
	}); err != nil {
		return fmt.Errorf("failed to update the status of the loan proposal: %w", err)
	}

	agreement := &models.LoanAgreement{
		LoanProposalUuid: proposalUuid,
		LoanRequestUuid:  requestUuid,
		LenderUuid:       loanProposal.UserUUID,
		BorrowerUuid:     acceptorUserUuid,
		Amount:           loanRequest.Amount,
		Interest:         loanProposal.MinInterest, // assured to max interest
		Duration:         loanProposal.MaxReturnDuration,
	}

	if _, err := l.loanAgreementDao.Upsert(ctx, agreement); err != nil {
		return fmt.Errorf("failed to upsert the loan agreement record: %w", err)
	}
	return nil
}

func (l loanProposalService) RejectGrant(ctx context.Context, rejectorUserUuid, proposalUuid, requestUuid string) error {
	loanRequest, err := l.loanRequestDao.FindById(ctx, requestUuid)
	if err != nil {
		return fmt.Errorf("failed to get the loan request: %w", err)
	}

	if loanRequest.UserUUID != rejectorUserUuid {
		return fmt.Errorf("cannot reject someone else's grant")
	}
	if utils.FromPtr(loanRequest.Status) != string(models.LOAN_REQUEST_GRANTED) {
		return fmt.Errorf("request not found to have a grant associated with it")
	}
	if utils.FromPtr(loanRequest.ProposalUuid) != proposalUuid {
		return fmt.Errorf("request grant not found to be associated with the provided proposal uuid")
	}

	_, err = l.loanRequestDao.Update(ctx, requestUuid, &models.LoanRequest{
		Status:       utils.ToPtr(""),
		ProposalUuid: utils.ToPtr(""),
	})
	if err != nil {
		return fmt.Errorf("failed to update the status of the loan request: %w", err)
	}

	if _, err := l.loanProposalDao.Update(ctx, proposalUuid, &models.LoanProposal{
		Status: string(models.LOAN_PROPOSAL_AVAILABLE),
	}); err != nil {
		return fmt.Errorf("failed to update the status of the loan proposal: %w", err)
	}

	return nil
}
