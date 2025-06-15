package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/oktopriima/deals/bootstrap/config"
	"github.com/oktopriima/deals/bootstrap/postgres"
	"github.com/oktopriima/deals/helper"
	"github.com/oktopriima/deals/models"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func main() {
	var err error
	// Initialize the application bootstrap
	container := dig.New()

	if err = container.Provide(func() config.AppConfig {
		return config.NewAppConfig()
	}); err != nil {
		panic(err)
	}

	// provide postgres connection
	if err = container.Provide(func(cfg config.AppConfig) postgres.DBInstance {
		return postgres.NewDatabaseInstance(cfg)
	}); err != nil {
		panic(err)
	}

	// run the migration
	if err = container.Invoke(func(dbInstance postgres.DBInstance) error {
		db := dbInstance.Database()

		if err := runSeeder(db); err != nil {
			return err
		}

		return nil
	}); err != nil {
		panic(err)
	}
}

func runSeeder(db *gorm.DB) error {
	if err := seedUser(db); err != nil {
		return err
	}

	return nil
}

func seedUser(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.User{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		fmt.Println("users already seeded. Skipping...")
		return nil
	}

	admin := models.User{
		Username: "admin",
		Password: helper.GeneratePassword("admin123"),
		IsAdmin:  true,
		Salary:   0,
	}

	if err := db.Create(&admin).Error; err != nil {
		return err
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i <= 100; i++ {
		u := models.User{
			Username: fmt.Sprintf("user%d", i),
			Password: helper.GeneratePassword(fmt.Sprintf("pass%d", i)),
			IsAdmin:  false,
			Salary:   float64(4000000 + rng.Intn(6000000)),
		}
		if err := db.Create(&u).Error; err != nil {
			return err
		}
	}

	return nil
}
