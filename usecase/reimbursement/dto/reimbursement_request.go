package dto

import (
	"github.com/oktopriima/deals/models"
	"time"
)

type ReimbursementRequest struct {
	UserId      int64     `json:"user_id" validate:"required"`
	Date        time.Time `json:"date" validate:"required"`
	Amount      float64   `json:"amount" validate:"required"`
	Description string    `json:"description"`
}

func (r *ReimbursementRequest) ToModel() *models.Reimbursement {
	return &models.Reimbursement{
		UserID:      r.UserId,
		Date:        r.Date,
		Amount:      r.Amount,
		Description: r.Description,
	}
}
