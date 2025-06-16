package dto

import (
	"github.com/oktopriima/deals/models"
	"time"
)

type overtimeResponse struct {
	UserId   int64     `json:"user_id"`
	Date     time.Time `json:"date"`
	Duration int64     `json:"duration"`
}

type OvertimeResponse interface {
	GetObject() *overtimeResponse
}

func (r *overtimeResponse) GetObject() *overtimeResponse {
	return r
}

func NewOvertimeResponse(overtime *models.Overtime) OvertimeResponse {
	return &overtimeResponse{
		UserId:   overtime.UserID,
		Date:     overtime.Timestamp,
		Duration: overtime.Duration,
	}
}
