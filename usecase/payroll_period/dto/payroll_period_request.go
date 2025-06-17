package dto

import (
	"github.com/oktopriima/deals/models"
	"time"
)

type PayrollPeriodRequest struct {
	StartDate time.Time `json:"start_date" time_format:"2006-01-02"  validate:"required"`
	EndDate   time.Time `json:"end_date" time_format:"2006-01-02" validate:"required"`
}

func (p *PayrollPeriodRequest) ToModel() *models.PayrollPeriod {
	return &models.PayrollPeriod{
		StartDate: p.StartDate,
		EndDate:   p.EndDate,
		Processed: false,
	}
}
