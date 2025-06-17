package models

import (
	"github.com/oktopriima/deals/lib/custom_middleware"
	"gorm.io/gorm"
)

type Payslip struct {
	ID                 int64   `gorm:"primaryKey"`
	UserID             int64   `gorm:"not null;index:user_period,unique"`
	PayrollPeriodId    int64   `gorm:"not null;index:user_period,unique"`
	BasePayment        float64 `gorm:"not null"`
	DailyPayment       float64 `gorm:"not null"`
	HourlyPayment      float64 `gorm:"not null"`
	TotalWorkingDays   float64 `gorm:"not null"`
	TotalAttendance    float64 `gorm:"not null"`
	DeductionAmount    float64 `gorm:"not null"`
	TotalHourOvertime  float64 `gorm:"not null"`
	OvertimePayment    float64 `gorm:"not null"`
	Reimbursements     float64 `gorm:"not null"`
	TotalPayment       float64 `gorm:"not null"`
	Users              string  `gorm:"type:jsonb"`
	ListAttendances    string  `gorm:"type:jsonb"`
	ListOvertimes      string  `gorm:"type:jsonb"`
	ListReimbursements string  `gorm:"type:jsonb"`
	BaseModel
}

func (p *Payslip) TableName() string {
	return "payslips"
}

func (p *Payslip) AfterCreate(tx *gorm.DB) error {
	return logAudit(tx, "create", p.TableName(), p.ID, p)
}

func (p *Payslip) BeforeCreate(tx *gorm.DB) error {
	ctx := tx.Statement.Context
	var userID *int64
	if uid, ok := ctx.Value(custom_middleware.UserId).(int64); ok {
		userID = &uid
	}

	p.CreatedBy = userID
	p.UpdatedBy = userID
	return nil
}

func (p *Payslip) BeforeUpdate(tx *gorm.DB) error {
	ctx := tx.Statement.Context
	var userID *int64
	if uid, ok := ctx.Value(custom_middleware.UserId).(int64); ok {
		userID = &uid
	}

	p.UpdatedBy = userID
	return nil
}

func (p *Payslip) AfterUpdate(tx *gorm.DB) error {
	return logAudit(tx, "update", p.TableName(), p.ID, p)
}
