package custom_middleware

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/oktopriima/deals/helper"
	jwthandle "github.com/oktopriima/deals/lib/jwtHandle"
	"net/http"
)

const (
	Token       = "TOKEN"
	AuthUser    = "AUTH_USER"
	AuthUserObj = "AUTH_USER_OBJ"
)

func Auth(token jwthandle.AccessToken) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request()
			oldCtx := c.Request().Context()

			headerToken, err := helper.HeaderExtractor("Authorization", r)
			if err != nil {
				return c.JSON(http.StatusForbidden, echo.Map{
					"code":    http.StatusForbidden,
					"message": err.Error(),
				})
			}

			if !token.Validate(headerToken) {
				return c.JSON(http.StatusForbidden, echo.Map{
					"code":    http.StatusForbidden,
					"message": "Invalid authorization token",
				})
			}

			e, err := jwthandle.Extract(headerToken, token.GetSignatureKey())
			if err != nil {
				return c.JSON(http.StatusForbidden, echo.Map{
					"code":    http.StatusForbidden,
					"message": err.Error(),
				})
			}

			obj, ok := e.Obj.(map[string]interface{})
			if !ok {
				return c.JSON(http.StatusForbidden, echo.Map{
					"code":    http.StatusForbidden,
					"message": "Invalid authorization object",
				})
			}

			ctx := context.WithValue(oldCtx, Token, headerToken)
			ctx = context.WithValue(ctx, AuthUser, e.Id)
			ctx = context.WithValue(ctx, AuthUserObj, obj)

			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}
