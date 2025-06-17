package payroll

import (
	"context"
	"fmt"
	"github.com/oktopriima/deals/lib/custom_middleware"
	"github.com/oktopriima/deals/usecase/payroll/dto"
)

func (p *payrollUsecase) FindByUserUsecase(ctx context.Context, periodId int64) (dto.DetailPayrollResponse, error) {
	userId, err := custom_middleware.GetAuthenticatedUser(ctx)
	if err != nil {
		return nil, err
	}

	payrollPeriod, err := p.payrollPeriodRepo.Find(ctx, periodId)
	if err != nil {
		return nil, err
	}

	payslip, err := p.payslipRepo.FindByUserUsecase(ctx, userId, periodId)
	if err != nil {
		return nil, err
	}

	output := dto.NewDetailPayrollResponse(payslip, payrollPeriod)
	if output == nil {
		return nil, fmt.Errorf("failed generate response")
	}

	return output, nil
}
