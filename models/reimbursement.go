package models

import (
	"github.com/oktopriima/deals/lib/custom_middleware"
	"gorm.io/gorm"
	"time"
)

type Reimbursement struct {
	ID          int64     `gorm:"primaryKey"`
	UserID      int64     `gorm:"not null;index"`
	Date        time.Time `gorm:"not null;type:timestamp without time zone"`
	Amount      float64   `gorm:"not null"`
	Description string
	BaseModel
}

func (r *Reimbursement) TableName() string {
	return "reimbursements"
}

func (r *Reimbursement) AfterCreate(tx *gorm.DB) error {
	return logAudit(tx, "create", r.TableName(), r.ID, r)
}

func (r *Reimbursement) BeforeCreate(tx *gorm.DB) error {
	ctx := tx.Statement.Context
	var userID *int64
	if uid, ok := ctx.Value(custom_middleware.UserId).(int64); ok {
		userID = &uid
	}

	r.CreatedBy = userID
	r.UpdatedBy = userID
	return nil
}

func (r *Reimbursement) BeforeUpdate(tx *gorm.DB) error {
	ctx := tx.Statement.Context
	var userID *int64
	if uid, ok := ctx.Value(custom_middleware.UserId).(int64); ok {
		userID = &uid
	}

	r.UpdatedBy = userID
	return nil
}

func (r *Reimbursement) AfterUpdate(tx *gorm.DB) error {
	return logAudit(tx, "update", r.TableName(), r.ID, r)
}
