package models

import "time"

type BaseModel struct {
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	CreatedBy *int64    `gorm:"index"`
	UpdatedBy *int64    `gorm:"index"`
}
