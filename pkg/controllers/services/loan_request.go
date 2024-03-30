package services

import (
	"bitbucket.com/finease/backend/pkg/utils"
	"context"
	"fmt"
	"time"

	// "time"

	"bitbucket.com/finease/backend/pkg/dao"
	"bitbucket.com/finease/backend/pkg/models"
	"github.com/google/uuid"
	// "github.com/google/uuid"
)

type LoanRequest interface {
	Create(ctx context.Context, creatorUserUuid string, loanRequest *models.LoanRequest) (*models.LoanRequest, error)
	Update(ctx context.Context, ownerUserUuid, id string, patch *models.LoanRequest) (*models.LoanRequest, error)
	Delete(ctx context.Context, ownerUserUuid, id string) error
	FindByUserId(ctx context.Context, ownerUserUuid string) ([]*models.LoanRequest, error)
	FindByProposalId(ctx context.Context, invokerUserUuid, proposalUuid string) ([]*models.LoanRequest, error)
	FindById(ctx context.Context, id string) (*models.LoanRequest, error)
	FindAll(ctx context.Context) ([]*models.LoanRequest, error)
	Accept(ctx context.Context, approverUserUuid, requestUuid string) error
	Reject(ctx context.Context, approverUserUuid, requestUuid string) error
}

type loanRequestService struct {
	userDao          dao.User
	loanRequestDao   dao.LoanRequest
	loanProposalDao  dao.LoanProposal
	loanAgreementDao dao.LoanAgreement
}

func NewLoanRequestService(loanRequestDao dao.LoanRequest,
	loanProposalDao dao.LoanProposal,
	loanAgreementDao dao.LoanAgreement,
	userDao dao.User) LoanRequest {
	return &loanRequestService{
		loanRequestDao:   loanRequestDao,
		loanProposalDao:  loanProposalDao,
		loanAgreementDao: loanAgreementDao,
		userDao:          userDao,
	}
}

func (l loanRequestService) Create(ctx context.Context, creatorUserUuid string, loanRequest *models.LoanRequest) (*models.LoanRequest, error) {
	loanRequest.CreatedAt, loanRequest.UpdatedAt = time.Now(), time.Now()
	loanRequest.Uuid = uuid.New().String()
	if loanRequest.Status == nil {
		loanRequest.Status = utils.ToPtr("")
	}

	if loanRequest.ProposalUuid == nil {
		return nil, fmt.Errorf("nil loan proposal uuid not allowed")
	}
	if utils.FromPtr(loanRequest.ProposalUuid) != "" {
		loanProposal, err := l.loanProposalDao.FindById(ctx, *loanRequest.ProposalUuid)
		if err != nil {
			return nil, fmt.Errorf("loan proposal not found with the provided uuid")
		}
		if loanProposal.UserUUID == creatorUserUuid {
			return nil, fmt.Errorf("cannot request a loan to yourself")
		}
		if loanRequest.MaxInterest != loanRequest.MinInterest {
			return nil, fmt.Errorf("min_interest and max_interest must match when creating a request against a proposal")
		}
	}

	createdLoanRequest, err := l.loanRequestDao.Create(ctx, loanRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to create the loan request: %w", err)
	}

	return createdLoanRequest, err
}

func (l loanRequestService) Update(ctx context.Context, ownerUserUuid, loanRequestUuid string, patch *models.LoanRequest) (*models.LoanRequest, error) {
	if patch.Uuid != "" {
		return nil, fmt.Errorf("not allowed to update UUID")
	}

	loanRequestFound, err := l.loanRequestDao.FindById(ctx, loanRequestUuid)
	if err != nil {
		return nil, fmt.Errorf("failed to find the loan request corresponding to the provided loan request uuid: %w", err)
	}

	if ownerUserUuid != loanRequestFound.UserUUID {
		return nil, fmt.Errorf("not authorized to update the loan request of some other user")
	}

	updatedLoanRequest, err := l.loanRequestDao.Update(ctx, loanRequestUuid, patch)
	if err != nil {
		return nil, fmt.Errorf("failed to update the loan request: %w", err)
	}
	return updatedLoanRequest, nil
}

func (l loanRequestService) Delete(ctx context.Context, ownerUserUuid, loanRequestUuid string) error {
	loanRequestFound, err := l.loanRequestDao.FindById(ctx, loanRequestUuid)
	if err != nil {
		return fmt.Errorf("failed to find the loan request corresponding to the provided loan request uuid: %w", err)
	}

	if ownerUserUuid != loanRequestFound.UserUUID {
		return fmt.Errorf("not authorized to delete the loan request of some other user")
	}

	if err := l.loanRequestDao.Delete(ctx, loanRequestUuid); err != nil {
		return fmt.Errorf("failed to delete the loan request: %w", err)
	}
	return nil
}

func (l loanRequestService) FindById(ctx context.Context, id string) (*models.LoanRequest, error) {
	loanRequests, err := l.loanRequestDao.FindById(ctx, id)
	if err != nil {
		return &models.LoanRequest{}, fmt.Errorf("failed to fetch your loan requests: %w", err)
	}
	return loanRequests, nil
}

func (l loanRequestService) FindByUserId(ctx context.Context, ownerUserUuid string) ([]*models.LoanRequest, error) {
	loanRequests, err := l.loanRequestDao.FindByUserId(ctx, ownerUserUuid)
	if err != nil {
		return []*models.LoanRequest{}, fmt.Errorf("failed to fetch your loan requests: %w", err)
	}
	return loanRequests, nil
}

func (l loanRequestService) FindByProposalId(ctx context.Context, invokerUserUuid string, proposalUuid string) ([]*models.LoanRequest, error) {
	proposal, err := l.loanProposalDao.FindById(ctx, proposalUuid)
	if err != nil {
		return []*models.LoanRequest{}, fmt.Errorf("failed to find the loan proposal by the provided id: %w", err)
	}

	if proposal.UserUUID != invokerUserUuid {
		return []*models.LoanRequest{}, fmt.Errorf("user not found to be the owner of the loan proposal")
	}

	loanRequests, err := l.loanRequestDao.FindByProposalId(ctx, proposalUuid)
	if err != nil {
		return []*models.LoanRequest{}, fmt.Errorf("failed to fetch your loan requests: %w", err)
	}
	return loanRequests, nil
}

func (l loanRequestService) FindAll(ctx context.Context) ([]*models.LoanRequest, error) {
	loanRequests, err := l.loanRequestDao.FindAll(ctx)
	if err != nil {
		return []*models.LoanRequest{}, fmt.Errorf("failed to fetch your loan requests: %w", err)
	}
	return loanRequests, nil
}

func (l loanRequestService) Accept(ctx context.Context, approverUserUuid, requestUuid string) error {
	return l.acceptOrReject(ctx, approverUserUuid, requestUuid, models.LOAN_REQUEST_ACCEPTED)
}

func (l loanRequestService) Reject(ctx context.Context, denyUserUuid, requestUuid string) error {
	return l.acceptOrReject(ctx, denyUserUuid, requestUuid, models.LOAN_REQUEST_REJECTED)
}

func (l loanRequestService) acceptOrReject(ctx context.Context, userUuid, requestUuid string, acceptOrReject models.LoanRequestStatus) error {
	loanRequest, err := l.loanRequestDao.FindById(ctx, requestUuid)
	if err != nil {
		return fmt.Errorf("failed to get the loan request: %w", err)
	}

	if utils.FromPtr(loanRequest.Status) != "" {
		if utils.FromPtr(loanRequest.Status) == string(acceptOrReject) {
			return nil
		}
		return fmt.Errorf("can't change the status after already approving/denying the loan request")
	}

	proposalUuid := utils.FromPtr(loanRequest.ProposalUuid)
	if proposalUuid == "" {
		return fmt.Errorf("not allowed to approve/deny this proposal uuid")
	}
	loanProposal, err := l.loanProposalDao.FindById(ctx, proposalUuid)
	if err != nil {
		return fmt.Errorf("failed to get the loan proposal's details: %w", err)
	}

	if userUuid != loanProposal.UserUUID {
		return fmt.Errorf("not allowed to approve request made to a proposal owned by someone else")
	}

	if _, err := l.loanRequestDao.Update(ctx, requestUuid, &models.LoanRequest{
		Status: utils.ToPtr(string(acceptOrReject)),
	}); err != nil {
		return fmt.Errorf("failed to update the loan request's status: %w", err)
	}

	if acceptOrReject == models.LOAN_REQUEST_REJECTED {
		if _, err := l.loanProposalDao.Update(ctx, proposalUuid, &models.LoanProposal{
			Status: string(models.LOAN_PROPOSAL_AVAILABLE),
		}); err != nil {
			return fmt.Errorf("failed to update the status of the loan proposal: %w", err)
		}
		return nil
	}

	if _, err := l.loanProposalDao.Update(ctx, proposalUuid, &models.LoanProposal{
		Status: string(models.LOAN_PROPOSAL_OFFERED),
	}); err != nil {
		return fmt.Errorf("failed to update the status of the loan proposal: %w", err)
	}

	// code path leading to loan request getting accepted
	agreement := &models.LoanAgreement{
		LoanProposalUuid: proposalUuid,
		LoanRequestUuid:  requestUuid,
		LenderUuid:       loanProposal.UserUUID,
		BorrowerUuid:     loanRequest.UserUUID,
		Amount:           loanRequest.Amount,
		Duration:         loanRequest.DurationToPay,
		Interest:         loanRequest.MinInterest, // assured to match max interest
	}

	if _, err := l.loanAgreementDao.Upsert(ctx, agreement); err != nil {
		return fmt.Errorf("failed to upsert the loan agreement record: %w", err)
	}

	return nil
}
