package models

import "time"

type Attendance struct {
	ID        int64     `gorm:"primaryKey"`
	UserID    int64     `gorm:"not null;index"`
	Timestamp time.Time `gorm:"not null"`
}

func (Attendance) TableName() string {
	return "attendances"
}
