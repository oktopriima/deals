package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/oktopriima/deals/usecase/payroll_period"
	"github.com/oktopriima/deals/usecase/payroll_period/dto"
	"net/http"
	"strconv"
)

type PayrollPeriodHandler struct {
	uc payroll_period.PayrollPeriodUsecase
}

func NewPayrollPeriodHandler(uc payroll_period.PayrollPeriodUsecase) *PayrollPeriodHandler {
	return &PayrollPeriodHandler{
		uc: uc,
	}
}

// ServeCreate handler for create endpoint
func (p *PayrollPeriodHandler) ServeCreate(ctx echo.Context) error {
	var req dto.PayrollPeriodRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	if err := ctx.Validate(req); err != nil {
		return err
	}

	output, err := p.uc.Create(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, echo.Map{
			"code":    http.StatusUnprocessableEntity,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, echo.Map{
		"code":    http.StatusCreated,
		"message": "success",
		"data":    output.GetObject(),
	})
}

// ServeFind handler for find endpoint
func (p *PayrollPeriodHandler) ServeFind(ctx echo.Context) error {
	idParam := ctx.Param("id")
	periodID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	output, err := p.uc.Find(ctx.Request().Context(), periodID)
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

// ServeList handler for list endpoint
func (p *PayrollPeriodHandler) ServeList(ctx echo.Context) error {
	periods, err := p.uc.List(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, echo.Map{
			"code":    http.StatusUnprocessableEntity,
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, echo.Map{
		"code":    http.StatusOK,
		"message": "success",
		"data":    periods.GetObject(),
	})
}
