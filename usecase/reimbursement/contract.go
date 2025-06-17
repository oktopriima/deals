package reimbursement

import (
	"context"
	"github.com/oktopriima/deals/repository"
	"github.com/oktopriima/deals/usecase/reimbursement/dto"
)

type reimbursementUsecase struct {
	reimbursementRepository repository.ReimbursementRepository
}

type ReimbursementUsecase interface {
	Store(req dto.ReimbursementRequest, ctx context.Context) (dto.ReimbursementResponse, error)
}

func NewReimbursementUsecase(reimbursementRepository repository.ReimbursementRepository) ReimbursementUsecase {
	return &reimbursementUsecase{
		reimbursementRepository: reimbursementRepository,
	}
}
