package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/oktopriima/deals/usecase/payroll"
	"github.com/oktopriima/deals/usecase/payroll/dto"
	"net/http"
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
