package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/oktopriima/deals/lib/custom_middleware"
	"github.com/oktopriima/deals/usecase/reimbursement"
	"github.com/oktopriima/deals/usecase/reimbursement/dto"
	"net/http"
)

type ReimbursementHandler struct {
	uc reimbursement.ReimbursementUsecase
}

func NewReimbursementHandler(usecase reimbursement.ReimbursementUsecase) *ReimbursementHandler {
	return &ReimbursementHandler{
		uc: usecase,
	}
}

func (r *ReimbursementHandler) Serve(ctx echo.Context) error {
	var req dto.ReimbursementRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	userId, err := custom_middleware.GetAuthenticatedUser(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{
			"code":    http.StatusUnauthorized,
			"message": err.Error(),
		})
	}
	req.UserId = userId

	if err := ctx.Validate(&req); err != nil {
		return err
	}

	output, err := r.uc.Store(req, ctx.Request().Context())
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
