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

	// Create the database
	for _, series := range data.SeriesData {
		db.Create(&series)
	}

	// Create view for all episodes in order of series, season, episode
	baseQuery := db.Model(&models.Episode{}).
		Select(`
			series.title AS series_title, 
			seasons.season_number AS season_number, 
			episodes.episode_number AS episode_number, 
			episodes.title AS episode_title, 
			episodes.air_date AS air_date, 
			episodes.star_date AS star_date
		`).
		Joins("JOIN seasons ON episodes.season_id = seasons.id").
		Joins("JOIN series ON seasons.series_id = series.id").
		Order("series.id, seasons.season_number, episodes.episode_number")

	db.Migrator().CreateView("episodes_view", gorm.ViewOption{
		Query:   baseQuery,
		Replace: true,
	})

	// Create view for all episodes in order of series, season, episode
	airDateQuery := db.Model(&models.Episode{}).
		Select(`
			series.title AS series_title, 
			seasons.season_number AS season_number, 
			episodes.episode_number AS episode_number, 
			episodes.title AS episode_title, 
			episodes.air_date AS air_date, 
			episodes.star_date AS star_date
		`).
		Joins("JOIN seasons ON episodes.season_id = seasons.id").
		Joins("JOIN series ON seasons.series_id = series.id").
		Order("air_date")

	db.Migrator().CreateView("episodes_air_date_view", gorm.ViewOption{
		Query:   airDateQuery,
		Replace: true,
	})

	// Create view for all episodes in order of series, season, episode
	starDateQuery := db.Model(&models.Episode{}).
		Select(`
			series.title AS series_title, 
			seasons.season_number AS season_number, 
			episodes.episode_number AS episode_number, 
			episodes.title AS episode_title, 
			episodes.air_date AS air_date, 
			episodes.star_date AS star_date
		`).
		Joins("JOIN seasons ON episodes.season_id = seasons.id").
		Joins("JOIN series ON seasons.series_id = series.id").
		Order("star_date")

	db.Migrator().CreateView("episodes_star_date_view", gorm.ViewOption{
		Query:   starDateQuery,
		Replace: true,
	})
}
