package dto

import "github.com/oktopriima/deals/models"

type runPayrollResponse struct {
	PayrollPeriodId int64  `json:"payroll_period_id"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
	IsProcessed     bool   `json:"is_processed"`
}

type RunPayrollResponse interface {
	GetObject() *runPayrollResponse
}

func (r *runPayrollResponse) GetObject() *runPayrollResponse {
	return r
}

func NewRunPayrollResponse(period *models.PayrollPeriod) *runPayrollResponse {
	return &runPayrollResponse{
		PayrollPeriodId: period.ID,
		StartDate:       period.StartDate.Format("2006-01-02"),
		EndDate:         period.EndDate.Format("2006-01-02"),
		IsProcessed:     period.Processed,
	}
}
