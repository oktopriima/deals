package models

type Payslip struct {
	ID             int64   `gorm:"primaryKey"`
	UserID         int64   `gorm:"not null;index:user_period,unique"`
	PayrollID      int64   `gorm:"not null;index:user_period,unique"`
	BasePay        float64 `gorm:"not null"`
	OvertimePay    float64 `gorm:"not null"`
	Reimbursements float64 `gorm:"not null"`
	TotalPay       float64 `gorm:"not null"`
}

func (Payslip) TableName() string {
	return "payslips"
}
