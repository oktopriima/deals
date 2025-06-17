package repository

import (
	"context"
	"errors"
	"github.com/oktopriima/deals/bootstrap/postgres"
	"github.com/oktopriima/deals/models"
	"gorm.io/gorm"
)

type reimbursementRepository struct {
	db *gorm.DB
}

type ReimbursementRepository interface {
	Store(reimbursement *models.Reimbursement, ctx context.Context) error
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
