package dto

import "github.com/oktopriima/deals/models"

type reimbursementResponse struct {
	UserId      int64   `json:"user_id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

type ReimbursementResponse interface {
	GetObject() *reimbursementResponse
}

func (r *reimbursementResponse) GetObject() *reimbursementResponse {
	return r
}

func NewReimbursementResponse(reimbursement *models.Reimbursement) ReimbursementResponse {
	return &reimbursementResponse{
		UserId:      reimbursement.UserID,
		Amount:      reimbursement.Amount,
		Description: reimbursement.Description,
	}
}
