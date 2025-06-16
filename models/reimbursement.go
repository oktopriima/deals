package models

import "time"

type Reimbursement struct {
	ID          int64     `gorm:"primaryKey"`
	UserID      int64     `gorm:"not null;index"`
	Date        time.Time `gorm:"not null;type:timestamp without time zone"`
	Amount      float64   `gorm:"not null"`
	Description string
}

func (Reimbursement) TableName() string {
	return "reimbursements"
}
