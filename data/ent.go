package data

import (
	"github.com/ericthomasca/star-trek-api/models"
)

var ENTData = models.Series{
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
}