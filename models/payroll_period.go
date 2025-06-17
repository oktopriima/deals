package models

import "time"

type PayrollPeriod struct {
	ID        int64     `gorm:"primaryKey"`
	StartDate time.Time `gorm:"not null;index:unique_period,unique;type:timestamp without time zone"`
	EndDate   time.Time `gorm:"not null;index:unique_period,unique;type:timestamp without time zone"`
	Processed bool      `gorm:"not null;default:false"`
	BaseModel
}

func (PayrollPeriod) TableName() string {
	return "payroll_periods"
}
