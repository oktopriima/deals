package dto

import (
	"time"

	"github.com/oktopriima/deals/models"
)

type attendanceResponse struct {
	UserId    int64     `json:"user_id"`
	Timestamp time.Time `json:"timestamp"`
	Date      string    `json:"date"`
	Time      string    `json:"time"`
	Zone      string    `json:"zone"`
}

type AttendanceResponse interface {
	GetObject() *attendanceResponse
}

func (r *attendanceResponse) GetObject() *attendanceResponse {
	return r
}

func NewAttendanceResponse(attendance *models.Attendance) AttendanceResponse {
	return &attendanceResponse{
		UserId:    attendance.UserID,
		Timestamp: attendance.Timestamp,
		Date:      attendance.Timestamp.Format("2006-01-02"),
		Time:      attendance.Timestamp.Format("15:04:05"),
		Zone:      attendance.Timestamp.Location().String(),
	}
}
