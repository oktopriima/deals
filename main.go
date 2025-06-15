package main

import (
	"github.com/oktopriima/deals/bootstrap"
	"github.com/oktopriima/deals/bootstrap/server"
	"github.com/oktopriima/deals/router"
)

func main() {
	c := bootstrap.NewBootstrap()
	err := c.Invoke(router.NewRouter)
	if err != nil {
		panic(err)
	}

	if err := c.Invoke(func(instance *server.EchoInstance) {
		instance.RunWithGracefullyShutdown()
	}); err != nil {
		panic(err)
	}
}
