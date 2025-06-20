package main

import (
	"github.com/oktopriima/deals/bootstrap/config"
	"github.com/oktopriima/deals/bootstrap/postgres"
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

		if err := runMigration(db); err != nil {
			return err
		}

		return nil
	}); err != nil {
		panic(err)
	}
}

func runMigration(db *gorm.DB) error {
	m := models.DB{DB: db}
	if err := m.AutoMigrate(); err != nil {
		return err
	}

	return nil
}
