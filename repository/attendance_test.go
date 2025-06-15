package repository

import (
	"context"
	"testing"
	"time"

	"github.com/oktopriima/deals/models"
	"github.com/stretchr/testify/assert"
)

// mockAttendance is a simple Attendance model for testing
func mockAttendance() *models.Attendance {
	return &models.Attendance{
		ID:         1,
		UserID:     1,
		Timestamp:  time.Now(),
		DateString: time.Now().Format("2006-01-02"),
		TimeString: time.Now().Format("15:04:05"),
	}
}

func TestAttendanceRepository_Store_Success(t *testing.T) {
	db := setupTestDB(t)
	repo := &attendanceRepository{db: db}
	ctx := context.Background()
	attendance := mockAttendance()

	err := repo.Store(attendance, ctx)
	assert.NoError(t, err)
}

func TestAttendanceRepository_Store_DBError(t *testing.T) {
	// Create a gorm.DB that will always return an error on Create
	repo := &attendanceRepository{db: nil}
	ctx := context.Background()
	attendance := mockAttendance()

	err := repo.Store(attendance, ctx)
	assert.Error(t, err)
}
func TestAttendanceRepository_CheckAlreadyExists_Found(t *testing.T) {
	db := setupTestDB(t)
	repo := &attendanceRepository{db: db}
	ctx := context.Background()
	attendance := mockAttendance()

	// Store the attendance first
	err := repo.Store(attendance, ctx)
	assert.NoError(t, err)

	exists := repo.CheckAlreadyExists(attendance.UserID, attendance.DateString, ctx)
	assert.True(t, exists)
}

func TestAttendanceRepository_CheckAlreadyExists_NotFound(t *testing.T) {
	db := setupTestDB(t)
	repo := &attendanceRepository{db: db}
	ctx := context.Background()

	exists := repo.CheckAlreadyExists(9999, "2099-01-01", ctx)
	assert.False(t, exists)
}

func TestAttendanceRepository_CheckAlreadyExists_DBError(t *testing.T) {
	repo := &attendanceRepository{db: nil}
	ctx := context.Background()

	exists := repo.CheckAlreadyExists(1, "2023-01-01", ctx)
	assert.False(t, exists)
}
