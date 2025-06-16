package models

import "time"

type Attendance struct {
	ID         int64     `gorm:"primaryKey"`
	UserID     int64     `gorm:"not null;index"`
	Timestamp  time.Time `gorm:"not null;type:timestamp without time zone"`
	DateString string    `gorm:"not null;index:date_user"`
	TimeString string    `gorm:"not null;index:time_user"`
}

func (Attendance) TableName() string {
	return "attendances"
}
