package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/ericthomasca/star-trek-api/controllers"
	"github.com/ericthomasca/star-trek-api/db"
	"github.com/ericthomasca/star-trek-api/seed"
)

func main() {
	// load db
	dbConn, err := db.LoadDB()
	if err != nil {
		log.Fatalf("Error loading database: %v", err)
	}

	// seed db
	seed.SeedDatabase(dbConn)

	// create router
	router := gin.Default()

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// add routes
	router.GET("/episodes", controllers.GetEpisodes)
	router.GET("/movies", controllers.GetMovies)

	// Start server
	router.Run(":61234") // listen and serve on 0.0.0.0:61234

}
