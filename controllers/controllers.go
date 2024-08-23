package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ericthomasca/star-trek-api/db"
	"github.com/ericthomasca/star-trek-api/models"
)

// GetEpisodes handles the /episodes route and returns episodes in JSON format
func GetEpisodes(c *gin.Context) {
	// Connect to the database
	dbConn := db.DB

	// Create an empty slice of episodes
	var episodes []models.EpisodeView

	// Query the database view
	err := dbConn.Table("episodes_view").Find(&episodes).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error retrieving episodes",
		})
		return
	}

	// Send the episodes as JSON response
	c.JSON(http.StatusOK, episodes)
}

// GetMovies handles the /movies route and returns movies in JSON format
func GetMovies(c *gin.Context) {
	// Connect to the database
	dbConn := db.DB	

	// Create an empty slice of movies
	var movies []models.Movie

	// Query the database view
	err := dbConn.Table("movies").Find(&movies).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error retrieving movies",
		})
		return
	}

	// Send the movies as JSON response
	c.JSON(http.StatusOK, movies)
}
