package dto

import (
	"time"

	"github.com/oktopriima/deals/models"
)

type AttendanceRequest struct {
	UserId    int64     `json:"user_id"`
	Timestamp time.Time `json:"timestamp"`
}

func (a AttendanceRequest) ToModel() *models.Attendance {
	return &models.Attendance{
		UserID:     a.UserId,
		Timestamp:  a.Timestamp,
		DateString: time.Now().Format("2006-01-02"),
		TimeString: time.Now().Format("15:04:05"),
	}
}
