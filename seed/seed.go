package seed

import (
	"github.com/ericthomasca/star-trek-api/data"
	"github.com/ericthomasca/star-trek-api/models"
	"gorm.io/gorm"
)

// SeedDatabase seeds the database with sample data.
func SeedDatabase(db *gorm.DB) {
	// Drop all tables
	db.Migrator().DropTable(&models.Series{}, &models.Season{}, &models.Episode{})

	// AutoMigrate will create the tables if they do not exist
	db.AutoMigrate(&models.Series{}, &models.Season{}, &models.Episode{})

	var seriesData = []models.Series{
		{
			Title: "Star Trek: Deep Space Nine",
			Seasons: []models.Season{
				{
					SeasonNumber: 1,
					Episodes:     []models.Episode{},
				},
				{
					SeasonNumber: 2,
					Episodes:     []models.Episode{},
				},
				{
					SeasonNumber: 3,
					Episodes:     []models.Episode{},
				},
				{
					SeasonNumber: 4,
					Episodes:     []models.Episode{},
				},
				{
					SeasonNumber: 5,
					Episodes:     []models.Episode{},
				},
				{
					SeasonNumber: 6,
					Episodes:     []models.Episode{},
				},
				{
					SeasonNumber: 7,
					Episodes:     []models.Episode{},
				},
			},
		},
		{
			Title: "Star Trek: Voyager",
			Seasons: []models.Season{
				{
					SeasonNumber: 1,
					Episodes:     []models.Episode{},
				},
				{
					SeasonNumber: 2,
					Episodes:     []models.Episode{},
				},
				{
					SeasonNumber: 3,
					Episodes:     []models.Episode{},
				},
				{
					SeasonNumber: 4,
					Episodes:     []models.Episode{},
				},
				{
					SeasonNumber: 5,
					Episodes:     []models.Episode{},
				},
				{
					SeasonNumber: 6,
					Episodes:     []models.Episode{},
				},
				{
					SeasonNumber: 7,
					Episodes:     []models.Episode{},
				},
			},
		},
		{
			Title: "Star Trek: Enterprise",
			Seasons: []models.Season{
				{
					SeasonNumber: 1,
					Episodes:     []models.Episode{},
				},
				{
					SeasonNumber: 2,
					Episodes:     []models.Episode{},
				},
				{
					SeasonNumber: 3,
					Episodes:     []models.Episode{},
				},
				{
					SeasonNumber: 4,
					Episodes:     []models.Episode{},
				},
			},
		},
	}

	db.Create(&data.TOSData)
	db.Create(&data.TASData)
	db.Create(&data.TNGData)

	// Seed series data
	for _, s := range seriesData {
		db.Create(&s)
	}
}

