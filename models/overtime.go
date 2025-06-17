package models

import (
	"github.com/oktopriima/deals/lib/custom_middleware"
	"gorm.io/gorm"
	"time"
)

type Overtime struct {
	ID         int64     `gorm:"primaryKey"`
	UserID     int64     `gorm:"not null;index"`
	Timestamp  time.Time `gorm:"not null;index:user_date;type:timestamp without time zone"`
	DateString string    `gorm:"not null;index:user_date"`
	Duration   int64     `gorm:"not null;check:duration >= 1 AND duration <= 3"`
	BaseModel
}

func (o *Overtime) TableName() string {
	return "overtimes"
}

func (o *Overtime) AfterCreate(tx *gorm.DB) error {
	return logAudit(tx, "create", o.TableName(), o.ID, o)
}

func (o *Overtime) BeforeCreate(tx *gorm.DB) error {
	ctx := tx.Statement.Context
	var userID *int64
	if uid, ok := ctx.Value(custom_middleware.UserId).(int64); ok {
		userID = &uid
	}

	o.CreatedBy = userID
	o.UpdatedBy = userID
	return nil
}

func (o *Overtime) BeforeUpdate(tx *gorm.DB) error {
	ctx := tx.Statement.Context
	var userID *int64
	if uid, ok := ctx.Value(custom_middleware.UserId).(int64); ok {
		userID = &uid
	}

	o.UpdatedBy = userID
	return nil
}

func (o *Overtime) AfterUpdate(tx *gorm.DB) error {
	return logAudit(tx, "update", o.TableName(), o.ID, o)
}
