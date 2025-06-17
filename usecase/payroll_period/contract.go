package payroll_period

import (
	"context"
	"github.com/oktopriima/deals/repository"
	"github.com/oktopriima/deals/usecase/payroll_period/dto"
)

type payrollPeriodUsecase struct {
	payrollPeriodRepository repository.PayrollPeriodRepository
}

type PayrollPeriodUsecase interface {
	Find(ctx context.Context, id int64) (dto.PayrollPeriodResponse, error)
	Create(ctx context.Context, request dto.PayrollPeriodRequest) (dto.PayrollPeriodResponse, error)
	List(ctx context.Context) (dto.ListPayrollPeriodResponse, error)
}

func NewPayrollPeriodUsecase(periodRepository repository.PayrollPeriodRepository) PayrollPeriodUsecase {
	return &payrollPeriodUsecase{
		payrollPeriodRepository: periodRepository,
	}
}
