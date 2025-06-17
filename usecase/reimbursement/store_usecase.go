package reimbursement

import (
	"context"
	"github.com/oktopriima/deals/usecase/reimbursement/dto"
)

func (r *reimbursementUsecase) Store(req dto.ReimbursementRequest, ctx context.Context) (dto.ReimbursementResponse, error) {
	if err := r.reimbursementRepository.Store(req.ToModel(), ctx); err != nil {
		return nil, err
	}

	return dto.NewReimbursementResponse(req.ToModel()), nil
}
