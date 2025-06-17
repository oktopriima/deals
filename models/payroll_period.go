package models

import (
	"github.com/oktopriima/deals/lib/custom_middleware"
	"gorm.io/gorm"
	"time"
)

type PayrollPeriod struct {
	ID        int64     `gorm:"primaryKey"`
	StartDate time.Time `gorm:"not null;index:unique_period,unique;type:timestamp without time zone"`
	EndDate   time.Time `gorm:"not null;index:unique_period,unique;type:timestamp without time zone"`
	Processed bool      `gorm:"not null;default:false"`
	BaseModel
}

func (p *PayrollPeriod) TableName() string {
	return "payroll_periods"
}

func (p *PayrollPeriod) AfterCreate(tx *gorm.DB) error {
	return logAudit(tx, "create", p.TableName(), p.ID, p)
}

func (p *PayrollPeriod) BeforeCreate(tx *gorm.DB) error {
	ctx := tx.Statement.Context
	var userID *int64
	if uid, ok := ctx.Value(custom_middleware.UserId).(int64); ok {
		userID = &uid
	}

	p.CreatedBy = userID
	p.UpdatedBy = userID
	return nil
}

func (p *PayrollPeriod) BeforeUpdate(tx *gorm.DB) error {
	ctx := tx.Statement.Context
	var userID *int64
	if uid, ok := ctx.Value(custom_middleware.UserId).(int64); ok {
		userID = &uid
	}

	p.UpdatedBy = userID
	return nil
}

func (p *PayrollPeriod) AfterUpdate(tx *gorm.DB) error {
	return logAudit(tx, "update", p.TableName(), p.ID, p)
}
