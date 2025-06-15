package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oktopriima/deals/usecase/authentication"
	"github.com/oktopriima/deals/usecase/authentication/dto"
)

type AuthenticationHandler struct {
	uc authentication.AuthenticationUsecase
}

func NewAuthenticationHandler(uc authentication.AuthenticationUsecase) *AuthenticationHandler {
	return &AuthenticationHandler{
		uc: uc,
	}
}

func (h *AuthenticationHandler) LoginByEmail(c echo.Context) error {
	var req dto.AuthenticationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	output, err := h.uc.LoginUsecase(req, c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, output)
}
