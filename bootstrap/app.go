package bootstrap

import "go.uber.org/dig"

func NewBootstrap() *dig.Container {
	c := dig.New()

	c = NewHttpServer(c)

	return c
}
