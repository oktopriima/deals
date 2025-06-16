package dto

import (
	"github.com/oktopriima/deals/models"
	"time"
)

type OvertimeRequest struct {
	UserId   int64
	Date     time.Time `json:"date" validate:"required"`
	Duration int64     `json:"duration" validate:"required,gt=0,lte=3"`
}

func (o *OvertimeRequest) ToModel() *models.Overtime {
	return &models.Overtime{
		UserID:     o.UserId,
		Timestamp:  o.Date,
		Duration:   o.Duration,
		DateString: o.Date.Format("2006-01-02"),
	}
}
