package data

import (
	"github.com/ericthomasca/star-trek-api/models"
)

// Create the TASData variable with the provided data
var TASData = models.Series{
	ID:    1, // Assuming an ID for the series
	Title: "Star Trek: The Animated Series",
	Seasons: []models.Season{
		{
			SeasonNumber: 1,
			Episodes: []models.Episode{
				{
					EpisodeNumber: 1,
					Title:         "Beyond the Farthest Star",
					AirDate:       strPtr("1973-09-08"),
					StarDate:      floatPtr(5521.3),
				},
				{
					EpisodeNumber: 2,
					Title:         "Yesteryear",
					AirDate:       strPtr("1973-09-15"),
					StarDate:      floatPtr(5373.4),
				},
				{
					EpisodeNumber: 3,
					Title:         "One of Our Planets is Missing",
					AirDate:       strPtr("1973-09-22"),
					StarDate:      floatPtr(5371.3),
				},
				{
					EpisodeNumber: 4,
					Title:         "The Lorelei Signal",
					AirDate:       strPtr("1973-09-29"),
					StarDate:      floatPtr(5483.7),
				},
				{
					EpisodeNumber: 5,
					Title:         "More Tribbles, More Troubles",
					AirDate:       strPtr("1973-10-06"),
					StarDate:      floatPtr(5392.4),
				},
				{
					EpisodeNumber: 6,
					Title:         "The Survivor",
					AirDate:       strPtr("1973-10-13"),
					StarDate:      floatPtr(5143.3),
				},
				{
					EpisodeNumber: 7,
					Title:         "The Infinite Vulcan",
					AirDate:       strPtr("1973-10-20"),
					StarDate:      floatPtr(5554.4),
				},
				{
					EpisodeNumber: 8,
					Title:         "The Magicks of Megas-tu",
					AirDate:       strPtr("1973-10-27"),
					StarDate:      floatPtr(1254.4),
				},
				{
					EpisodeNumber: 9,
					Title:         "Once Upon a Planet",
					AirDate:       strPtr("1973-11-03"),
					StarDate:      floatPtr(5591.2),
				},
				{
					EpisodeNumber: 10,
					Title:         "Mudd's Passion",
					AirDate:       strPtr("1973-11-10"),
					StarDate:      floatPtr(4978.5),
				},
				{
					EpisodeNumber: 11,
					Title:         "The Terratin Incident",
					AirDate:       strPtr("1973-11-17"),
					StarDate:      floatPtr(5577.3),
				},
				{
					EpisodeNumber: 12,
					Title:         "The Time Trap",
					AirDate:       strPtr("1973-11-24"),
					StarDate:      floatPtr(5267.2),
				},
				{
					EpisodeNumber: 13,
					Title:         "The Ambergris Element",
					AirDate:       strPtr("1973-12-01"),
					StarDate:      floatPtr(5499.9),
				},
				{
					EpisodeNumber: 14,
					Title:         "The Slaver Weapon",
					AirDate:       strPtr("1973-12-15"),
					StarDate:      floatPtr(4187.3),
				},
				{
					EpisodeNumber: 15,
					Title:         "The Eye of the Beholder",
					AirDate:       strPtr("1974-01-05"),
					StarDate:      floatPtr(5501.2),
				},
				{
					EpisodeNumber: 16,
					Title:         "The Jihad",
					AirDate:       strPtr("1974-01-12"),
					StarDate:      floatPtr(5683.1),
				},
			},
		},
		{
			SeasonNumber: 2,
			Episodes: []models.Episode{
				{
					EpisodeNumber: 1,
					Title:         "The Pirates of Orion",
					AirDate:       strPtr("1974-09-07"),
					StarDate:      floatPtr(6334.1),
				},
				{
					EpisodeNumber: 2,
					Title:         "Bem",
					AirDate:       strPtr("1974-09-17"),
					StarDate:      floatPtr(7403.6),
				},
				{
					EpisodeNumber: 3,
					Title:         "The Practical Joker",
					AirDate:       strPtr("1974-09-21"),
					StarDate:      floatPtr(3183.3),
				},
				{
					EpisodeNumber: 4,
					Title:         "Albatross",
					AirDate:       strPtr("1974-09-28"),
					StarDate:      floatPtr(5275.6),
				},
				{
					EpisodeNumber: 5,
					Title:         "How Sharper Than a Serpent's Tooth",
					AirDate:       strPtr("1974-10-05"),
					StarDate:      floatPtr(6063.4),
				},
				{
					EpisodeNumber: 6,
					Title:         "The Counter-Clock Incident",
					AirDate:       strPtr("1974-10-12"),
					StarDate:      floatPtr(6770.3),
				},
			},
		},
	},
}
