package repository

import (
	"context"
	"errors"

	"github.com/oktopriima/deals/bootstrap/postgres"
	"github.com/oktopriima/deals/models"
	"gorm.io/gorm"
)

type attendanceRepository struct {
	db *gorm.DB
}

type AttendanceRepository interface {
	Store(m *models.Attendance, c context.Context) error
	CheckAlreadyExists(userID int64, dateString string, c context.Context) bool
}

func NewAttendanceRepository(dbInstance postgres.DBInstance) AttendanceRepository {
	return &attendanceRepository{db: dbInstance.Database()}
}

func (a *attendanceRepository) Store(m *models.Attendance, c context.Context) error {
	if a.db == nil {
		return errors.New("attendance repository not initialized")
	}

	db := a.db.WithContext(c)
	result := db.Create(&m)

	return result.Error
}

// CheckAlreadyExists implements AttendanceRepository.
func (a *attendanceRepository) CheckAlreadyExists(userID int64, dateString string, c context.Context) bool {
	var count int64
	if a.db == nil {
		return false
	}

	db := a.db.WithContext(c)
	result := db.Model(&models.Attendance{}).
		Where("user_id = ? AND date_string = ?", userID, dateString).
		Count(&count)

	if result.Error != nil {
		return false
	}

	return count > 0
}
