package dto

import (
	"time"

	"github.com/oktopriima/deals/models"
)

type AttendanceRequest struct {
	UserId    int64     `json:"user_id" validate:"required,gt=0"`
	Timestamp time.Time `json:"timestamp"  validate:"required"`
}

func (a AttendanceRequest) ToModel() *models.Attendance {
	return &models.Attendance{
		UserID:     a.UserId,
		Timestamp:  a.Timestamp,
		DateString: a.Timestamp.Format("2006-01-02"),
		TimeString: a.Timestamp.Format("15:04:05"),
	}
}
