package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/oktopriima/deals/handler"
)

func NewRouter(
	e *echo.Echo,
	authHandler *handler.AuthenticationHandler,
) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	route := e.Group("/api")

	{
		route.GET("/ping", func(c echo.Context) error {
			return c.JSON(200, map[string]string{
				"message": "pong!!!",
			})
		})
	}

	{
		authRoute := route.Group("/auth")
		authRoute.POST("", authHandler.LoginByEmail)
	}
}
