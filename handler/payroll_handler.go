package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/oktopriima/deals/usecase/payroll"
	"github.com/oktopriima/deals/usecase/payroll/dto"
	"net/http"
	"strconv"
)

type PayrollHandler struct {
	uc payroll.PayrollUsecase
}

func NewPayrollHandler(usecase payroll.PayrollUsecase) *PayrollHandler {
	return &PayrollHandler{
		uc: usecase,
	}
}

func (p *PayrollHandler) Serve(ctx echo.Context) error {
	var req dto.RunPayrollRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	if err := ctx.Validate(req); err != nil {
		return err
	}

	output, err := p.uc.RunUsecase(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, echo.Map{
			"code":    http.StatusUnprocessableEntity,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"code":    http.StatusOK,
		"message": "ok",
		"data":    output.GetObject(),
	})
}

func (p *PayrollHandler) ServeFindByUser(ctx echo.Context) error {
	strId := ctx.Param("payrollPeriodId")
	if strId == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"code":    http.StatusBadRequest,
			"message": "Please provide payroll period id",
		})
	}

	periodId, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	output, err := p.uc.FindByUserUsecase(ctx.Request().Context(), periodId)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, echo.Map{
			"code":    http.StatusUnprocessableEntity,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"code":    http.StatusOK,
		"message": "ok",
		"data":    output.GetObject(),
	})
}
