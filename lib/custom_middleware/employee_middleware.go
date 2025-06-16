package custom_middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Employee() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			ctx := c.Request().Context()

			user, ok := ctx.Value(AuthUserObj).(map[string]interface{})
			if !ok {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"code":    http.StatusUnauthorized,
					"message": "Failed to retrieve user information",
				})
			}

			isAdmin, ok := user["IsAdmin"].(bool)
			if !ok || isAdmin == true {
				return c.JSON(http.StatusForbidden, echo.Map{
					"code":    http.StatusForbidden,
					"message": "You don't have permission to view this resource",
				})
			}

			return next(c)
		}
	}
}
