package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/oktopriima/deals/handler"
	"github.com/oktopriima/deals/lib/custom_middleware"
	"github.com/oktopriima/deals/lib/jwtHandle"
)

func NewRouter(
	e *echo.Echo,
	jwt jwthandle.AccessToken,
	authHandler *handler.AuthenticationHandler,
	adminAttendanceHandler *handler.AdminAttendanceHandler,
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

	{
		adminRoute := route.Group("/admin")
		adminRoute.Use(custom_middleware.Auth(jwt))
		adminRoute.Use(custom_middleware.Admin())

		{
			adminAttendanceRoute := adminRoute.Group("/attendance")
			adminAttendanceRoute.POST("", adminAttendanceHandler.Serve)
		}
	}
}
