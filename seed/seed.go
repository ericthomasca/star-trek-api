package seed

import (
	"gorm.io/gorm"
)

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

// SeedDatabase seeds the database with sample data.
func SeedDatabase(db *gorm.DB) {
	// Drop all tables
	db.Migrator().DropTable(&Series{}, &Season{}, &Episode{})

	// AutoMigrate will create the tables if they do not exist
	db.AutoMigrate(&Series{}, &Season{}, &Episode{})

	var seriesData = []Series{
		{
			Title: "Star Trek: The Original Series",
			Seasons: []Season{
				{
					SeasonNumber: 1,
					Episodes: []Episode{
						{
							EpisodeNumber: 1,
							Title:         "The Man Trap",
							AirDate:       strPtr("1966-09-08"),
							StarDate:      floatPtr(1513.1),
						},
						{
							EpisodeNumber: 2,
							Title:         "Charlie X",
							AirDate:       strPtr("1966-09-15"),
							StarDate:      floatPtr(1533.6),
						},
						{
							EpisodeNumber: 3,
							Title:         "Where No Man Has Gone Before",
							AirDate:       strPtr("1966-09-22"),
							StarDate:      floatPtr(1312.4),
						},
						{
							EpisodeNumber: 4,
							Title:         "The Naked Time",
							AirDate:       strPtr("1966-09-29"),
							StarDate:      floatPtr(1704.2),
						},
						{
							EpisodeNumber: 5,
							Title:         "The Enemy Within",
							AirDate:       strPtr("1966-10-06"),
							StarDate:      floatPtr(1672.1),
						},
						{
							EpisodeNumber: 6,
							Title:         "Mudd's Women",
							AirDate:       strPtr("1966-10-13"),
							StarDate:      floatPtr(1329.8),
						},
						{
							EpisodeNumber: 7,
							Title:         "What Are Little Girls Made Of?",
							AirDate:       strPtr("1966-10-20"),
							StarDate:      floatPtr(2712.4),
						},
						{
							EpisodeNumber: 8,
							Title:         "Miri",
							AirDate:       strPtr("1966-10-27"),
							StarDate:      floatPtr(2713.5),
						},
						{
							EpisodeNumber: 9,
							Title:         "Dagger of the Mind",
							AirDate:       strPtr("1966-11-03"),
							StarDate:      floatPtr(2715.1),
						},
						{
							EpisodeNumber: 10,
							Title:         "The Corbomite Maneuver",
							AirDate:       strPtr("1966-11-10"),
							StarDate:      floatPtr(1512.2),
						},
						{
							EpisodeNumber: 11,
							Title:         "The Menagerie, Part I",
							AirDate:       strPtr("1966-11-17"),
							StarDate:      floatPtr(3012.4),
						},
						{
							EpisodeNumber: 12,
							Title:         "The Menagerie, Part II",
							AirDate:       strPtr("1966-12-08"),
							StarDate:      floatPtr(3013.1),
						},
						{
							EpisodeNumber: 13,
							Title:         "The Conscience of the King",
							AirDate:       strPtr("1966-12-01"),
							StarDate:      floatPtr(2817.6),
						},
						{
							EpisodeNumber: 14,
							Title:         "Balance of Terror",
							AirDate:       strPtr("1966-12-15"),
							StarDate:      floatPtr(1709.2),
						},
						{
							EpisodeNumber: 15,
							Title:         "Shore Leave",
							AirDate:       strPtr("1966-12-29"),
							StarDate:      floatPtr(3025.3),
						},
						{
							EpisodeNumber: 16,
							Title:         "The Galileo Seven",
							AirDate:       strPtr("1967-01-05"),
							StarDate:      floatPtr(2821.5),
						},
						{
							EpisodeNumber: 17,
							Title:         "The Squire of Gothos",
							AirDate:       strPtr("1967-01-12"),
							StarDate:      floatPtr(2124.5),
						},
						{
							EpisodeNumber: 18,
							Title:         "Arena",
							AirDate:       strPtr("1967-01-19"),
							StarDate:      floatPtr(3045.6),
						},
						{
							EpisodeNumber: 19,
							Title:         "Tomorrow Is Yesterday",
							AirDate:       strPtr("1967-01-26"),
							StarDate:      floatPtr(3113.2),
						},
						{
							EpisodeNumber: 20,
							Title:         "Court Martial",
							AirDate:       strPtr("1967-02-02"),
							StarDate:      floatPtr(2947.3),
						},
						{
							EpisodeNumber: 21,
							Title:         "The Return of the Archons",
							AirDate:       strPtr("1967-02-09"),
							StarDate:      floatPtr(3156.2),
						},
						{
							EpisodeNumber: 22,
							Title:         "Space Seed",
							AirDate:       strPtr("1967-02-16"),
							StarDate:      floatPtr(3141.9),
						},
						{
							EpisodeNumber: 23,
							Title:         "A Taste of Armageddon",
							AirDate:       strPtr("1967-02-23"),
							StarDate:      floatPtr(3192.1),
						},
						{
							EpisodeNumber: 24,
							Title:         "This Side of Paradise",
							AirDate:       strPtr("1967-03-02"),
							StarDate:      floatPtr(3417.3),
						},
						{
							EpisodeNumber: 25,
							Title:         "The Devil in the Dark",
							AirDate:       strPtr("1967-03-09"),
							StarDate:      floatPtr(3196.1),
						},
						{
							EpisodeNumber: 26,
							Title:         "Errand of Mercy",
							AirDate:       strPtr("1967-03-23"),
							StarDate:      floatPtr(3198.4),
						},
						{
							EpisodeNumber: 27,
							Title:         "The Alternative Factor",
							AirDate:       strPtr("1967-03-30"),
							StarDate:      floatPtr(3087.6),
						},
						{
							EpisodeNumber: 28,
							Title:         "The City on the Edge of Forever",
							AirDate:       strPtr("1967-04-06"),
							StarDate:      nil, // No StarDate
						},
						{
							EpisodeNumber: 29,
							Title:         "Operation: Annihilate!",
							AirDate:       strPtr("1967-04-13"),
							StarDate:      floatPtr(3287.2),
						},
					},
				},
				{
					SeasonNumber: 2,
					Episodes:     []Episode{},
				},
				{
					SeasonNumber: 3,
					Episodes:     []Episode{},
				},
			},
		},
		{
			Title: "Star Trek: The Animated Series",
			Seasons: []Season{
				{
					SeasonNumber: 1,
					Episodes:     []Episode{},
				},
				{
					SeasonNumber: 2,
					Episodes:     []Episode{},
				},
			},
		},
		{
			Title: "Star Trek: The Next Generation",
			Seasons: []Season{
				{
					SeasonNumber: 1,
					Episodes:     []Episode{},
				},
				{
					SeasonNumber: 2,
					Episodes:     []Episode{},
				},
				{
					SeasonNumber: 3,
					Episodes:     []Episode{},
				},
				{
					SeasonNumber: 4,
					Episodes:     []Episode{},
				},
				{
					SeasonNumber: 5,
					Episodes:     []Episode{},
				},
				{
					SeasonNumber: 6,
					Episodes:     []Episode{},
				},
				{
					SeasonNumber: 7,
					Episodes:     []Episode{},
				},
			},
		},
		{
			Title: "Star Trek: Deep Space Nine",
			Seasons: []Season{
				{
					SeasonNumber: 1,
					Episodes:     []Episode{},
				},
				{
					SeasonNumber: 2,
					Episodes:     []Episode{},
				},
				{
					SeasonNumber: 3,
					Episodes:     []Episode{},
				},
				{
					SeasonNumber: 4,
					Episodes:     []Episode{},
				},
				{
					SeasonNumber: 5,
					Episodes:     []Episode{},
				},
				{
					SeasonNumber: 6,
					Episodes:     []Episode{},
				},
				{
					SeasonNumber: 7,
					Episodes:     []Episode{},
				},
			},
		},
		{
			Title: "Star Trek: Voyager",
			Seasons: []Season{
				{
					SeasonNumber: 1,
					Episodes:     []Episode{},
				},
				{
					SeasonNumber: 2,
					Episodes:     []Episode{},
				},
				{
					SeasonNumber: 3,
					Episodes:     []Episode{},
				},
				{
					SeasonNumber: 4,
					Episodes:     []Episode{},
				},
				{
					SeasonNumber: 5,
					Episodes:     []Episode{},
				},
				{
					SeasonNumber: 6,
					Episodes:     []Episode{},
				},
				{
					SeasonNumber: 7,
					Episodes:     []Episode{},
				},
			},
		},
		{
			Title: "Star Trek: Enterprise",
			Seasons: []Season{
				{
					SeasonNumber: 1,
					Episodes:     []Episode{},
				},
				{
					SeasonNumber: 2,
					Episodes:     []Episode{},
				},
				{
					SeasonNumber: 3,
					Episodes:     []Episode{},
				},
				{
					SeasonNumber: 4,
					Episodes:     []Episode{},
				},
			},
		},
	}

	// Seed series data
	for _, s := range seriesData {
		db.Create(&s)
	}
}

func strPtr(s string) *string {
	return &s
}

func floatPtr(f float64) *float64 {
	return &f
}
