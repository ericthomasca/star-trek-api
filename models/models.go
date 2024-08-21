package models

// Series represents a Star Trek series.
type Series struct {
	ID      uint     `gorm:"primaryKey"`      // The ID of the series
	Title   string   `gorm:"unique;not null"` // The title of the series
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