package bootstrap

import (
	"github.com/oktopriima/deals/repository"
	"go.uber.org/dig"
)

func NewRepository(c *dig.Container) *dig.Container {
	var err error

	if err = c.Provide(repository.NewUserRepository); err != nil {
		panic(err)
	}

	if err = c.Provide(repository.NewAttendanceRepository); err != nil {
		panic(err)
	}

	if err = c.Provide(repository.NewOvertimeRepository); err != nil {
		panic(err)
	}

	if err = c.Provide(repository.NewReimbursementRepository); err != nil {
		panic(err)
	}

	if err = c.Provide(repository.NewPayrollPeriodRepository); err != nil {
		panic(err)
	}

	return c
}
