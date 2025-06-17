package models

import (
	"encoding/json"
	"github.com/oktopriima/deals/lib/custom_middleware"
	"gorm.io/gorm"
	"time"
)

type AuditLog struct {
	ID        int64   `gorm:"primaryKey"`
	Action    string  `gorm:"not null"`
	Model     string  `gorm:"not null"`
	ModelID   int64   `gorm:"not null"`
	UserID    *int64  `gorm:"index"`
	Data      string  `gorm:"type:jsonb"`
	RequestID *string `gorm:"index"`
	IPAddress *string
	Timestamp time.Time `gorm:"autoCreateTime"`
}

func logAudit(tx *gorm.DB, action, model string, modelID int64, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	ctx := tx.Statement.Context
	var userID *int64
	if uid, ok := ctx.Value(custom_middleware.UserId).(int64); ok {
		userID = &uid
	}

	var requestID *string
	if rid, ok := ctx.Value(custom_middleware.RequestId).(string); ok {
		requestID = &rid
	}

	var ipAddress *string
	if ip, ok := ctx.Value(custom_middleware.IpAddress).(string); ok {
		ipAddress = &ip
	}

	audit := AuditLog{
		Action:    action,
		Model:     model,
		ModelID:   modelID,
		UserID:    userID,
		Data:      string(data),
		RequestID: requestID,
		IPAddress: ipAddress,
		Timestamp: time.Now(),
	}
	return tx.Create(&audit).Error
}
