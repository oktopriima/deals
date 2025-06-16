package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/oktopriima/deals/usecase/attendance"
	"github.com/oktopriima/deals/usecase/attendance/dto"
	"net/http"
)

type AdminAttendanceHandler struct {
	uc attendance.AttendanceUsecase
}

func NewAdminAttendanceHandler(uc attendance.AttendanceUsecase) *AdminAttendanceHandler {
	return &AdminAttendanceHandler{
		uc: uc,
	}
}

func (a *AdminAttendanceHandler) Serve(ctx echo.Context) error {
	var req dto.AttendanceRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

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
