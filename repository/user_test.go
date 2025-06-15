package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/oktopriima/deals/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// helper to create a test user in the DB
func seedUser(db *gorm.DB, username string) *models.User {
	user := &models.User{
		Username: username,
	}
	db.Create(user)
	return user
}

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	err = db.AutoMigrate(
		&models.User{},
		&models.Attendance{},
		&models.Overtime{},
		&models.Reimbursement{},
		&models.PayrollPeriod{},
		&models.Payslip{},
	)
	assert.NoError(t, err)
	return db
}

type mockDBInstance struct {
	db *gorm.DB
}

// Close implements postgres.DBInstance.
func (m *mockDBInstance) Close() {
	db, err := m.db.DB()
	if err == nil {
		db.Close()
	}
}

func (m *mockDBInstance) Database() *gorm.DB {
	return m.db
}

func TestFindByUsername_Success(t *testing.T) {
	db := setupTestDB(t)
	username := "testuser"
	seedUser(db, username)

	repo := NewUserRepository(&mockDBInstance{db: db})

	ctx := context.Background()
	user, err := repo.FindByUsername(username, ctx)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, username, user.Username)
}

func TestFindByUsername_NotFound(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(&mockDBInstance{db: db})

	ctx := context.Background()
	user, err := repo.FindByUsername("nonexistent", ctx)
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}
