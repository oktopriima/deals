package models

import "time"

type Attendance struct {
	ID           int64     `gorm:"primaryKey"`
	UserID       int64     `gorm:"not null;index"`
	TimestampIn  time.Time `gorm:"not null"`
	TimestampOut time.Time
}

func (Attendance) TableName() string {
	return "attendances"
}
