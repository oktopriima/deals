package dto

type RunPayrollRequest struct {
	PayrollPeriodId int64 `json:"payroll_period_id" validate:"required,number"`
}
