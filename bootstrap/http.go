package bootstrap

import (
	"github.com/labstack/echo/v4"
	"github.com/oktopriima/deals/bootstrap/config"
	"github.com/oktopriima/deals/bootstrap/postgres"
	"github.com/oktopriima/deals/bootstrap/server"
	"go.uber.org/dig"
)

func NewHttpServer(container *dig.Container) *dig.Container {
	var err error

	// provide config
	if err = container.Provide(func() config.AppConfig {
		return config.NewAppConfig()
	}); err != nil {
		panic(err)
	}

	// provide postgres connection
	if err = container.Provide(func(cfg config.AppConfig) postgres.DBInstance {
		return postgres.NewDatabaseInstance(cfg)
	}); err != nil {
		panic(err)
	}

	// provide echo instance
	if err = container.Provide(server.NewEchoInstance); err != nil {
		panic(err)
	}

	// provide router
	if err = container.Provide(func() *echo.Echo {
		e := echo.New()
		return e
	}); err != nil {
		panic(err)
	}

	return container
}
