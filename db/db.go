package db

import (
	"fmt"
	"github.com/icchy-san/uctest/config"
	"github.com/icchy-san/uctest/db/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

//go:generate mockgen -source db.go -destination db_mock.go -package db
type DB interface {
	GetInvoices(tx *gorm.DB, period_start_at, period_end_at *time.Time) ([]model.Invoice, error)
}

type db struct {
	client *gorm.DB
}

// initDB returns db client by connection
func initDB(env *config.Env) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		env.DBHost,
		env.DBUser,
		env.DBPassword,
		env.DBName,
	)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if migrateErr := database.AutoMigrate(
		model.Invoice{},
		model.Company{},
		model.BusinessPartner{},
		model.Bank{},
		model.User{},
	); migrateErr != nil {
		return nil, migrateErr
	}

	return database, nil
}

// return initialized db
func New(env *config.Env) (DB, error) {
	client, err := initDB(env)
	if err != nil {
		return nil, err
	}

	return &db{
		client: client,
	}, nil
}
