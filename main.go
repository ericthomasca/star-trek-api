package main

import (
	"log"

	"github.com/ericthomasca/star-trek-api/config"
	"github.com/ericthomasca/star-trek-api/seed"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
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

	seed.SeedDatabase(db)
}
