package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/oktopriima/deals/lib/custom_middleware"
	"github.com/oktopriima/deals/usecase/attendance"
	"github.com/oktopriima/deals/usecase/attendance/dto"
	"net/http"
)

type AttendanceHandler struct {
	uc attendance.AttendanceUsecase
}

func NewAttendanceHandler(uc attendance.AttendanceUsecase) *AttendanceHandler {
	return &AttendanceHandler{
		uc: uc,
	}
}

func (a *AttendanceHandler) Serve(ctx echo.Context) error {
	var req dto.AttendanceRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	// get user ID from token
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

	output, err := a.uc.CreateUsecase(req, ctx.Request().Context())
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
