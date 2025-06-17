package repository

import (
	"context"
	"errors"
	"github.com/oktopriima/deals/bootstrap/postgres"
	"github.com/oktopriima/deals/models"
	"gorm.io/gorm"
)

type payrollPeriodRepository struct {
	db *gorm.DB
}

type PayrollPeriodRepository interface {
	Store(ctx context.Context, period *models.PayrollPeriod) error
	Find(ctx context.Context, id int64) (*models.PayrollPeriod, error)
	List(ctx context.Context) ([]*models.PayrollPeriod, error)
	Update(ctx context.Context, period *models.PayrollPeriod) error
}

func NewPayrollPeriodRepository(instance postgres.DBInstance) PayrollPeriodRepository {
	return &payrollPeriodRepository{db: instance.Database()}
}

func (p *payrollPeriodRepository) Store(ctx context.Context, period *models.PayrollPeriod) error {
	if p.db == nil {
		return errors.New("payroll period repository not initialized")
	}

	db := p.db.WithContext(ctx)
	result := db.Create(&period)

	return result.Error
}

func (p *payrollPeriodRepository) Find(ctx context.Context, id int64) (*models.PayrollPeriod, error) {
	if p.db == nil {
		return nil, errors.New("payroll period repository not initialized")
	}

	var period models.PayrollPeriod
	result := p.db.WithContext(ctx).First(&period, id)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &period, nil
}

func (p *payrollPeriodRepository) List(ctx context.Context) ([]*models.PayrollPeriod, error) {
	if p.db == nil {
		return nil, errors.New("payroll period repository not initialized")
	}

	var periods []*models.PayrollPeriod
	result := p.db.WithContext(ctx).Find(&periods)
	if err := result.Error; err != nil {
		return nil, err
	}
	return periods, nil
}

func (p *payrollPeriodRepository) Update(ctx context.Context, period *models.PayrollPeriod) error {
	if p.db == nil {
		return errors.New("payroll period repository not initialized")
	}

	db := p.db.WithContext(ctx)
	result := db.Save(&period)

	return result.Error
}
