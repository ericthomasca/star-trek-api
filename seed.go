package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Series represents a Star Trek series.
type Series struct {
	ID       uint      `gorm:"primaryKey"` // The ID of the series
	Title    string    `gorm:"unique;not null"` // The title of the series
	Episodes []Episode // The episodes of the series
}

// Episode represents an episode in a Star Trek series.
type Episode struct {
	ID            uint     `gorm:"primaryKey"` // The ID of the episode
	SeriesID      uint     // The ID of the associated Series
	Series        Series   // The associated Series
	SeasonNumber  int      // The season number of the episode
	EpisodeNumber int      // The episode number within the season
	Title         string   // The title of the episode
	AirDate       *string  // The air date of the episode (nullable)
	StarDate      *float64 // The Stardate of the episode (nullable)
}

// SeedDatabase seeds the database with sample data.
func SeedDatabase() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// get database credentials from .env
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// connect to database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Drop all tables
	db.Migrator().DropTable(&Series{}, &Episode{})

	// AutoMigrate will create the tables if they do not exist
	db.AutoMigrate(&Series{}, &Episode{})

	// Read the JSON file
	seriesData, err := os.ReadFile("data/series.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Unmarshal the JSON data into a slice of Series
	var series []Series
	err = json.Unmarshal(seriesData, &series)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Seed series data
	for _, s := range series {
		db.Create(&s)
	}

}
