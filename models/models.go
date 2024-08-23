package models

// Series represents a Star Trek series.
type Series struct {
	ID      uint     `gorm:"primaryKey"` // The ID of the series
	Title   string   `gorm:"not null"`   // The title of the series
	Seasons []Season // The episodes of the series
}

// Season represents a season of a Star Trek series.
type Season struct {
	ID           uint `gorm:"primaryKey"` // The ID of the season
	SeriesID     uint // The ID of the associated Series
	SeasonNumber int  `gorm:"index"` // The season number
	Episodes     []Episode
}

// Episode represents an episode of a season in a Star Trek series.
type Episode struct {
	ID            uint     `gorm:"primaryKey"` // The ID of the episode
	SeasonID      uint     // The ID of the associated Series
	EpisodeNumber int      `gorm:"index"`    // The episode number within the season
	Title         string   `gorm:"not null"` // The title of the episode
	AirDate       *string  `gorm:"index"`    // The air date of the episode (nullable)
	StarDate      *float64 `gorm:"index"`    // The Stardate of the episode (nullable)
}

// Movie represents a Star Trek movie.
type Movie struct {
	ID          uint    `gorm:"primaryKey"` // The ID of the movie
	Title       string  `gorm:"not null"`   // The title of the movie
	ReleaseDate *string `gorm:"index"`      // The release date of the movie (nullable)
}

// EpisodeView represents the structure of the data in the episodes_view.
type EpisodeView struct {
	SeriesTitle    string  `json:"series_title"`
	SeasonNumber   int     `json:"season_number"`
	EpisodeNumber  int     `json:"episode_number"`
	EpisodeTitle   string  `json:"episode_title"`
	AirDate        *string `json:"air_date"`
	StarDate       *float64 `json:"star_date"`
}
