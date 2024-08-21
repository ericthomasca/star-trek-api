package data

import (
	"github.com/ericthomasca/star-trek-api/models"
)

var TASData = models.Series{
	Title: "Star Trek: The Animated Series",
	Seasons: []models.Season{
		{
			SeasonNumber: 1,
			Episodes: []models.Episode{},
		},
		{
			SeasonNumber: 2,
			Episodes:     []models.Episode{},
		},
	},
}