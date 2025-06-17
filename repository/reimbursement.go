package repository

import (
	"context"
	"errors"
	"github.com/oktopriima/deals/bootstrap/postgres"
	"github.com/oktopriima/deals/models"
	"gorm.io/gorm"
	"time"
)

type reimbursementRepository struct {
	db *gorm.DB
}

type ReimbursementRepository interface {
	Store(reimbursement *models.Reimbursement, ctx context.Context) error
	ListByUserID(userID int64, startDate, endDate time.Time, ctx context.Context) ([]*models.Reimbursement, error)
}

func NewReimbursementRepository(instance postgres.DBInstance) ReimbursementRepository {
	return &reimbursementRepository{db: instance.Database()}
}

func (r *reimbursementRepository) Store(reimbursement *models.Reimbursement, ctx context.Context) error {
	if r.db == nil {
		return errors.New("reimbursement repository not initialized")
	}

	db := r.db.WithContext(ctx)
	result := db.Create(&reimbursement)

	return result.Error
}

func (r *reimbursementRepository) ListByUserID(userID int64, startDate, endDate time.Time, ctx context.Context) ([]*models.Reimbursement, error) {
	if r.db == nil {
		return nil, errors.New("reimbursement repository not initialized")
	}

	endDate = endDate.Add(time.Second * 86399)
	db := r.db.WithContext(ctx)
	var reimbursements []*models.Reimbursement

	result := db.Where("user_id = ? AND date BETWEEN ? AND ?", userID, startDate.Format(time.RFC3339), endDate.Format(time.RFC3339)).
		Find(&reimbursements)

	return reimbursements, result.Error
}
