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

	for _, series := range data.SeriesData {
		db.Create(&series)
	}
}
