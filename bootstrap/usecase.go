package bootstrap

import (
	"github.com/oktopriima/deals/usecase/attendance"
	"github.com/oktopriima/deals/usecase/authentication"
	"github.com/oktopriima/deals/usecase/overtime"
	"github.com/oktopriima/deals/usecase/payroll_period"
	"github.com/oktopriima/deals/usecase/reimbursement"
	"go.uber.org/dig"
)

func NewUsecase(c *dig.Container) *dig.Container {
	var err error

	if err = c.Provide(authentication.NewAuthenticationUsecase); err != nil {
		panic(err)
	}

	if err = c.Provide(attendance.NewAttendanceUsecase); err != nil {
		panic(err)
	}

	if err = c.Provide(overtime.NewOvertimeUsecase); err != nil {
		panic(err)
	}

	if err = c.Provide(reimbursement.NewReimbursementUsecase); err != nil {
		panic(err)
	}

	if err = c.Provide(payroll_period.NewPayrollPeriodUsecase); err != nil {
		panic(err)
	}

	return c
}
