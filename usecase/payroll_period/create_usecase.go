package payroll_period

import (
	"context"
	"github.com/oktopriima/deals/usecase/payroll_period/dto"
)

func (p *payrollPeriodUsecase) Create(ctx context.Context, request dto.PayrollPeriodRequest) (dto.PayrollPeriodResponse, error) {
	if err := p.payrollPeriodRepository.Store(ctx, request.ToModel()); err != nil {
		return nil, err
	}

	return dto.NewPayrollPeriodResponse(request.ToModel()), nil
}
