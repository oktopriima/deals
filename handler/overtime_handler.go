package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/oktopriima/deals/lib/custom_middleware"
	"github.com/oktopriima/deals/usecase/overtime"
	"github.com/oktopriima/deals/usecase/overtime/dto"
	"net/http"
)

type OvertimeHandler struct {
	uc overtime.OvertimeUsecase
}

func NewOvertimeHandler(uc overtime.OvertimeUsecase) *OvertimeHandler {
	return &OvertimeHandler{uc: uc}
}

func (o *OvertimeHandler) Serve(ctx echo.Context) error {
	var req dto.OvertimeRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	userId, err := custom_middleware.GetAuthenticatedUser(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}
	req.UserId = userId

	if err := ctx.Validate(&req); err != nil {
		return err
	}

	output, err := o.uc.CreateUsecase(req, ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, echo.Map{
			"code":    http.StatusUnprocessableEntity,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"code":    http.StatusOK,
		"message": "success",
		"data":    output.GetObject(),
	})
}
