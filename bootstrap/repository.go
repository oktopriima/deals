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

	return c
}
