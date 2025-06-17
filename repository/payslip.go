package repository

import (
	"context"
	"errors"
	"github.com/oktopriima/deals/bootstrap/postgres"
	"github.com/oktopriima/deals/models"
	"gorm.io/gorm"
)

type payslipRepository struct {
	db *gorm.DB
}

type PayslipRepository interface {
	Store(ctx context.Context, payslip *models.Payslip) error
	FindByUserUsecase(ctx context.Context, userId, periodId int64) (*models.Payslip, error)
}

func NewPayslipRepository(instance postgres.DBInstance) PayslipRepository {
	return &payslipRepository{
		db: instance.Database(),
	}
}

func (p *payslipRepository) Store(ctx context.Context, payslip *models.Payslip) error {
	if p.db == nil {
		return errors.New("payslip repository is not initialized")
	}

	result := p.db.WithContext(ctx).Create(payslip)
	return result.Error
}

func (p *payslipRepository) FindByUserUsecase(ctx context.Context, userId, periodId int64) (*models.Payslip, error) {
	if p.db == nil {
		return nil, errors.New("payslip repository is not initialized")
	}

	var payslip *models.Payslip
	result := p.db.WithContext(ctx).Where("user_id = ? AND payroll_period_id = ?", userId, periodId).First(&payslip)
	return payslip, result.Error
}
