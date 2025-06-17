package repository

import (
	"context"
	"errors"
	"github.com/oktopriima/deals/bootstrap/postgres"
	"github.com/oktopriima/deals/models"
	"gorm.io/gorm"
	"time"
)

type overtimeRepository struct {
	db *gorm.DB
}

type OvertimeRepository interface {
	Store(overtime *models.Overtime, ctx context.Context) error
	FindByUserDate(userId int64, date time.Time, ctx context.Context) (*models.Overtime, error)
	Update(overtime *models.Overtime, ctx context.Context) error
	ListOvertimeByUserId(userId int64, startDate, endDate time.Time, ctx context.Context) ([]*models.Overtime, error)
}

func NewOvertimeRepository(instance postgres.DBInstance) OvertimeRepository {
	return &overtimeRepository{db: instance.Database()}
}

func (o *overtimeRepository) Store(overtime *models.Overtime, ctx context.Context) error {
	if o.db == nil {
		return errors.New("overtime repository not initialized")
	}

	db := o.db.WithContext(ctx)
	result := db.Create(&overtime)

	return result.Error
}

func (o *overtimeRepository) FindByUserDate(userId int64, date time.Time, ctx context.Context) (*models.Overtime, error) {
	if o.db == nil {
		return nil, errors.New("overtime repository not initialized")
	}

	db := o.db.WithContext(ctx)
	var overtime models.Overtime

	result := db.Where("user_id = ? AND date_string = ?", userId, date.Format("2006-01-02")).First(&overtime)

	if result.Error != nil {
		return nil, result.Error
	}

	return &overtime, nil
}

func (o *overtimeRepository) Update(overtime *models.Overtime, ctx context.Context) error {
	if o.db == nil {
		return errors.New("overtime repository not initialized")
	}

	db := o.db.WithContext(ctx)

	result := db.Save(&overtime)
	return result.Error
}

func (o *overtimeRepository) ListOvertimeByUserId(userId int64, startDate, endDate time.Time, ctx context.Context) ([]*models.Overtime, error) {
	if o.db == nil {
		return nil, errors.New("overtime repository not initialized")
	}

	endDate = endDate.Add(time.Second * 86399)
	var list []*models.Overtime
	db := o.db.WithContext(ctx)
	result := db.Where("user_id = ? AND timestamp BETWEEN ? AND ?", userId, startDate.Format(time.RFC3339), endDate.Format(time.RFC3339)).
		Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}

	return list, nil
}
