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

	// Seed the database with show data
	db.Create(&data.TOSData) // Star Trek: The Original Series
	db.Create(&data.TASData) // Star Trek: The Animated Series
	db.Create(&data.TNGData) // Star Trek: The Next Generation
	db.Create(&data.DS9Data) // Star Trek: Deep Space Nine
	db.Create(&data.VOYData) // Star Trek: Voyager
	db.Create(&data.ENTData) // Star Trek: Enterprise
}
