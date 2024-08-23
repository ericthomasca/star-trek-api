package db

import (
	"log"

	"github.com/ericthomasca/star-trek-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadDB() (*gorm.DB, error) {
	// load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// create database connection string
	dsn := cfg.GetDSN()

	// connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db

	return db, nil
}
