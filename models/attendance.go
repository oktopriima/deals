package models

import (
	"github.com/oktopriima/deals/lib/custom_middleware"
	"gorm.io/gorm"
	"time"
)

type Attendance struct {
	ID         int64     `gorm:"primaryKey"`
	UserID     int64     `gorm:"not null;index"`
	Timestamp  time.Time `gorm:"not null;type:timestamp without time zone"`
	DateString string    `gorm:"not null;index:date_user"`
	TimeString string    `gorm:"not null;index:time_user"`
	BaseModel
}

func (a *Attendance) TableName() string {
	return "attendances"
}

func (a *Attendance) BeforeCreate(tx *gorm.DB) error {
	ctx := tx.Statement.Context
	var userID *int64
	if uid, ok := ctx.Value(custom_middleware.UserId).(int64); ok {
		userID = &uid
	}

	a.CreatedBy = userID
	a.UpdatedBy = userID
	return nil
}

func (a *Attendance) AfterCreate(tx *gorm.DB) error {
	return logAudit(tx, "create", a.TableName(), a.ID, a)
}

func (a *Attendance) BeforeUpdate(tx *gorm.DB) error {
	ctx := tx.Statement.Context
	var userID *int64
	if uid, ok := ctx.Value(custom_middleware.UserId).(int64); ok {
		userID = &uid
	}

	a.UpdatedBy = userID
	return nil
}

func (a *Attendance) AfterUpdate(tx *gorm.DB) error {
	return logAudit(tx, "update", a.TableName(), a.ID, a)
}
