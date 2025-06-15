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

	return c
}
