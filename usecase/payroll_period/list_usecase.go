package payroll_period

import (
	"context"
	"github.com/oktopriima/deals/usecase/payroll_period/dto"
)

func (p *payrollPeriodUsecase) List(ctx context.Context) (dto.ListPayrollPeriodResponse, error) {
	periods, err := p.payrollPeriodRepository.List(ctx)
	if err != nil {
		return nil, err
	}

	return dto.NewListPayrollPeriodResponse(periods), nil
}
