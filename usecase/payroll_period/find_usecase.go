package payroll_period

import (
	"context"
	"github.com/oktopriima/deals/usecase/payroll_period/dto"
)

func (p *payrollPeriodUsecase) Find(ctx context.Context, id int64) (dto.PayrollPeriodResponse, error) {
	model, err := p.payrollPeriodRepository.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	return dto.NewPayrollPeriodResponse(model), nil
}
