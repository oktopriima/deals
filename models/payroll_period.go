package models

import "time"

type PayrollPeriod struct {
	ID        int64     `gorm:"primaryKey"`
	StartDate time.Time `gorm:"not null;index:unique_period,unique"`
	EndDate   time.Time `gorm:"not null;index:unique_period,unique"`
	Processed bool      `gorm:"not null;default:false"`
}

func (PayrollPeriod) TableName() string {
	return "payroll_periods"
}
