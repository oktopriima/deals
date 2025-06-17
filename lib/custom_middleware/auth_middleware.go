package custom_middleware

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/oktopriima/deals/helper"
	jwthandle "github.com/oktopriima/deals/lib/jwtHandle"
	"net/http"
	"strconv"
)

const (
	Token       = "TOKEN"
	AuthUser    = "AUTH_USER"
	AuthUserObj = "AUTH_USER_OBJ"
	UserId      = "USER_ID"
	IpAddress   = "IP_ADDRESS"
	RequestId   = "REQUEST_ID"
)

func Auth(token jwthandle.AccessToken) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request()
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

			uid, _ := strconv.ParseInt(e.Id, 10, 64)

			ctx := c.Request().Context()
			ctx = context.WithValue(ctx, Token, headerToken)
			ctx = context.WithValue(ctx, AuthUser, e.Id)
			ctx = context.WithValue(ctx, AuthUserObj, obj)
			ctx = context.WithValue(ctx, UserId, uid)
			ctx = context.WithValue(ctx, IpAddress, c.RealIP())
			ctx = context.WithValue(ctx, RequestId, c.Response().Header().Get(echo.HeaderXRequestID))

			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}
