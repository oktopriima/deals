package bootstrap

import (
	"github.com/oktopriima/deals/handler"
	"go.uber.org/dig"
)

func NewHandler(c *dig.Container) *dig.Container {
	var err error

	if err = c.Provide(handler.NewAuthenticationHandler); err != nil {
		panic(err)
	}

	if err = c.Provide(handler.NewAdminAttendanceHandler); err != nil {
		panic(err)
	}

	if err = c.Provide(handler.NewAttendanceHandler); err != nil {
		panic(err)
	}

	if err = c.Provide(handler.NewOvertimeHandler); err != nil {
		panic(err)
	}

	if err = c.Provide(handler.NewReimbursementHandler); err != nil {
		panic(err)
	}

	return c
}
