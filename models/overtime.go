package models

import "time"

type Overtime struct {
	ID         int64     `gorm:"primaryKey"`
	UserID     int64     `gorm:"not null;index"`
	Timestamp  time.Time `gorm:"not null;index:user_date;type:timestamp without time zone"`
	DateString string    `gorm:"not null;index:user_date"`
	Duration   int64     `gorm:"not null;check:hours >= 1 AND hours <= 3"`
}

func (Overtime) TableName() string {
	return "overtimes"
}
