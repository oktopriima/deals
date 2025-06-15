package bootstrap

import (
	"github.com/oktopriima/deals/usecase/authentication"
	"go.uber.org/dig"
)

func NewUsecase(c *dig.Container) *dig.Container {
	var err error

	if err = c.Provide(authentication.NewAuthenticationUsecase); err != nil {
		panic(err)
	}

	return c
}
