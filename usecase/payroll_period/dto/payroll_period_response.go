package dto

import "github.com/oktopriima/deals/models"

// single response
type payrollPeriodResponse struct {
	Id        int64  `json:"id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type PayrollPeriodResponse interface {
	GetObject() *payrollPeriodResponse
}

func (p *payrollPeriodResponse) GetObject() *payrollPeriodResponse {
	return p
}

func NewPayrollPeriodResponse(period *models.PayrollPeriod) PayrollPeriodResponse {
	return &payrollPeriodResponse{
		Id:        period.ID,
		StartDate: period.StartDate.Format("2006-01-02"),
		EndDate:   period.EndDate.Format("2006-01-02"),
	}
}

type listPayrollPeriodResponse struct {
	List []payrollPeriodResponse `json:"list"`
}

type ListPayrollPeriodResponse interface {
	GetObject() *listPayrollPeriodResponse
}

func (p *listPayrollPeriodResponse) GetObject() *listPayrollPeriodResponse {
	return p
}

func NewListPayrollPeriodResponse(periods []*models.PayrollPeriod) ListPayrollPeriodResponse {
	res := make([]payrollPeriodResponse, len(periods))
	for i, period := range periods {
		res[i] = payrollPeriodResponse{
			Id:        period.ID,
			StartDate: period.StartDate.Format("2006-01-02"),
			EndDate:   period.EndDate.Format("2006-01-02"),
		}
	}

	return &listPayrollPeriodResponse{
		List: res,
	}
}
