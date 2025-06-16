package models

import "time"

type Overtime struct {
	ID     int64     `gorm:"primaryKey"`
	UserID int64     `gorm:"not null;index"`
	Date   time.Time `gorm:"not null;index:user_date,unique;type:timestamp without time zone"`
	Hours  int       `gorm:"not null;check:hours >= 1 AND hours <= 3"`
}

func (Overtime) TableName() string {
	return "overtimes"
}
