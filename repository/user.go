package repository

import (
	"context"
	"errors"

	"github.com/oktopriima/deals/bootstrap/postgres"
	"github.com/oktopriima/deals/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	FindByUsername(email string, ctx context.Context) (*models.User, error)
	ListEmployees(ctx context.Context) ([]*models.User, error)
}

func NewUserRepository(dbInstance postgres.DBInstance) UserRepository {
	return &userRepository{
		db: dbInstance.Database(),
	}
}

// FindByUsername implements UserRepository.
func (u *userRepository) FindByUsername(email string, ctx context.Context) (*models.User, error) {
	var user models.User
	result := u.db.WithContext(ctx).Where("username = ?", email).First(&user)

	if result.Error != nil {
		return nil, result.Error

	}

	return &user, nil
}

func (u *userRepository) ListEmployees(ctx context.Context) ([]*models.User, error) {
	if u.db == nil {
		return nil, errors.New("db instance not initialized")
	}

	var users []*models.User
	result := u.db.WithContext(ctx).
		Where("is_admin = ?", false).
		Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
