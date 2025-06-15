package models

type Payslip struct {
	ID              int64   `gorm:"primaryKey"`
	UserID          int64   `gorm:"not null;index:user_period,unique"`
	PayrollPeriodId int64   `gorm:"not null;index:user_period,unique"`
	BasePayment     float64 `gorm:"not null"`
	OvertimePayment float64 `gorm:"not null"`
	Reimbursements  float64 `gorm:"not null"`
	TotalPayment    float64 `gorm:"not null"`
}

func (Payslip) TableName() string {
	return "payslips"
}
