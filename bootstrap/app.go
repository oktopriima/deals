package bootstrap

import "go.uber.org/dig"

func NewBootstrap() *dig.Container {
	c := dig.New()

	c = NewHttpServer(c)
	c = NewRepository(c)
	c = NewUsecase(c)
	c = NewHandler(c)

	return c
}
