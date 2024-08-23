package data

import (
	"github.com/ericthomasca/star-trek-api/models"
)

var MovieData = []models.Movie{
	{
		Title:       "Star Trek: The Motion Picture",
		ReleaseDate: strPtr("1979-12-07"),
	},
	{
		Title:       "Star Trek: The Wrath of Khan",
		ReleaseDate: strPtr("1982-06-04"),
	},
	{
		Title:       "Star Trek: The Search for Spock",
		ReleaseDate: strPtr("1984-06-01"),
	},
	{
		Title:       "Star Trek: The Voyage Home",
		ReleaseDate: strPtr("1986-11-26"),
	},
	{
		Title:       "Star Trek: The Final Frontier",
		ReleaseDate: strPtr("1989-06-09"),
	},
	{
		Title:       "Star Trek: The Undiscovered Country",
		ReleaseDate: strPtr("1991-12-06"),
	},
	{
		Title:       "Star Trek: Generations",
		ReleaseDate: strPtr("1994-11-18"),
	},
	{
		Title:       "Star Trek: First Contact",
		ReleaseDate: strPtr("1996-11-22"),
	},
	{
		Title:       "Star Trek: Insurrection",
		ReleaseDate: strPtr("1998-12-11"),
	},
	{
		Title:       "Star Trek: Nemesis",
		ReleaseDate: strPtr("2002-12-13"),
	},
}
