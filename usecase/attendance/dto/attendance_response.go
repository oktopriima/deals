package dto

import (
	"time"

	"github.com/oktopriima/deals/models"
)

type attendanceResponse struct {
	UserId    int64     `json:"user_id"`
	Timestamp time.Time `json:"timestamp"`
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
	}
}
