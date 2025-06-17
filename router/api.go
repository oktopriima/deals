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
	attendanceHandler *handler.AttendanceHandler,
	overtimeHandler *handler.OvertimeHandler,
	reimbursementHandler *handler.ReimbursementHandler,
	payrollPeriodHandler *handler.PayrollPeriodHandler,
	adminAttendanceHandler *handler.AdminAttendanceHandler,
) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())

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
		authRoute.POST("", authHandler.Serve)
	}

	{
		employeesRoute := route.Group("/employees")
		employeesRoute.Use(custom_middleware.Auth(jwt))
		employeesRoute.Use(custom_middleware.Employee())

		{
			attendanceRoute := employeesRoute.Group("/attendance")
			attendanceRoute.POST("", attendanceHandler.Serve)
		}

		{
			overtimeRoute := employeesRoute.Group("/overtime")
			overtimeRoute.POST("", overtimeHandler.Serve)
		}

		{
			reimbursementRoute := employeesRoute.Group("/reimbursement")
			reimbursementRoute.POST("", reimbursementHandler.Serve)
		}

		{
			payrollPeriodRoute := employeesRoute.Group("/payroll/period")
			payrollPeriodRoute.GET("/:id", payrollPeriodHandler.ServeFind)
			payrollPeriodRoute.GET("", payrollPeriodHandler.ServeList)
		}
	}

	{
		adminRoute := route.Group("/admin")
		adminRoute.Use(custom_middleware.Auth(jwt))
		adminRoute.Use(custom_middleware.Admin())

		{
			adminAttendanceRoute := adminRoute.Group("/attendance")
			adminAttendanceRoute.POST("", adminAttendanceHandler.Serve)
		}

		{
			adminPayrollPeriodRoute := adminRoute.Group("/payroll/period")
			adminPayrollPeriodRoute.POST("", payrollPeriodHandler.ServeCreate)
			adminPayrollPeriodRoute.GET("", payrollPeriodHandler.ServeList)
		}
	}
}
