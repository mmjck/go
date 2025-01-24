package config

import (
	"gopportunities/schemas"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgresql() (*gorm.DB, error) {

	logger = NewLogger("sqlite")
	dbPath := "./db/main.db"
	_, err := os.Stat(dbPath)

	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5434 sslmode=disable TimeZone=UTC"

	// Create DB and connect

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Errorf("postgresql opening error: %v", err)
		return nil, err
	}
	// Migrate the Schema
	err = db.AutoMigrate(&schemas.Opening{})
	logger.Info("AutoMigrate executed")
	if err != nil {
		logger.Errorf("postgresql automigration error: %v", err)
		return nil, err
	}
	// Return the DB
	return db, nil

}
