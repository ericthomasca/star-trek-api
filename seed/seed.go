package seed

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
	ID       uint      `gorm:"primaryKey"`                         // The ID of the series
	Title    string    `gorm:"unique;not null;column:title;index"` // The title of the series
	Episodes []Episode `gorm:"foreignKey:SeriesID"`                // The episodes of the series
}

// Episode represents an episode in a Star Trek series.
type Episode struct {
	ID            uint     `gorm:"primaryKey"`    // The ID of the episode
	SeriesID      uint     `gorm:"foreignKey:ID"` // The ID of the associated Series
	Series        Series   // The associated Series
	SeasonNumber  int      `gorm:"column:season_number"`  // The season number of the episode
	EpisodeNumber int      `gorm:"column:episode_number"` // The episode number within the season
	Title         string   `gorm:"column:title"`          // The title of the episode
	AirDate       *string  `gorm:"column:air_date"`       // The air date of the episode (nullable)
	StarDate      *float64 `gorm:"column:star_date"`      // The Stardate of the episode (nullable)
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

	// create database connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)

	// connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Drop all tables
	db.Migrator().DropTable(&Series{}, &Episode{})

	// AutoMigrate will create the tables if they do not exist
	db.AutoMigrate(&Series{}, &Episode{})

	// seed series data
	seedSeriesListFromJSON(db, "data/series.json")

	// seed tos season data
	// seedEpisodesFromDirectory(db, "data/seasons/tos")
}

// seedSeriesListFromJSON parses a JSON file and returns a slice of Series.
func seedSeriesListFromJSON(db *gorm.DB, path string) {
	// Read the JSON file for series
	seriesData, err := os.ReadFile(path)
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

// // seedEpisodeListFromJSON parses a JSON file and returns a slice of Episode.
// func seedEpisodeListFromJSON(db *gorm.DB, path string) {
// 	// Read the JSON file for episodes
// 	episodeData, err := os.ReadFile(path)
// 	if err != nil {
// 		log.Fatalf("Error reading file: %v", err)
// 	}

// 	// Unmarshal the JSON data into a slice of Episode
// 	var episodes []Episode
// 	err = json.Unmarshal(episodeData, &episodes)
// 	if err != nil {
// 		log.Fatalf("Error unmarshalling JSON: %v", err)
// 	}

// 	// Seed episode data
// 	for _, e := range episodes {
// 		db.Create(&e)
// 	}
// }

// // seedEpisodesFromDirectory parses a directory of JSON files and returns a slice of Episode.
// func seedEpisodesFromDirectory(db *gorm.DB, dirPath string) {
// 	files, err := os.ReadDir(dirPath)
// 	if err != nil {
// 		log.Fatalf("Error reading directory %s: %v", dirPath, err)
// 	}

// 	for _, file := range files {
// 		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
// 			filePath := fmt.Sprintf("%s/%s", dirPath, file.Name())
// 			fmt.Printf("Processing file: %s\n", filePath)
// 			seedEpisodeListFromJSON(db, filePath)
// 		}
// 	}
// }
