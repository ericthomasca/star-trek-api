Great! Given that your project is hosted on GitHub at github.com/ericthomaca/star-trek-api, here’s how you can structure and implement the data seeding based on your project’s structure and requirements.
Directory Structure and Files

To align with your project, you’ll need to:

    Create Data Files for Series and Episodes.
    Set Up Data Seeding to populate your database.

Here’s a step-by-step guide:
1. Define Series and Episodes in Go Files

series.go

Create a file to define the series data:

go

// series.go
package main

type Series struct {
    ID       uint      `gorm:"primaryKey"`
    Title    string    `gorm:"unique;not null;column:title;index"`
    Episodes []Episode `gorm:"foreignKey:SeriesID"`
}

// Define the series data
var seriesList = []Series{
    {
        ID:    1,
        Title: "The Original Series",
    },
    // Add more series if needed
}

seasons/tos/season1.go

Define episodes for Season 1 of "The Original Series":

go

// seasons/tos/season1.go
package tos

// Define episodes for Season 1
var Season1Episodes = []Episode{
    {
        SeriesID:      1,
        SeasonNumber:  1,
        EpisodeNumber: 1,
        Title:         "The Man Trap",
        AirDate:       ptrString("1966-09-08"),
        StarDate:      ptrFloat64(1513.1),
    },
    {
        SeriesID:      1,
        SeasonNumber:  1,
        EpisodeNumber: 2,
        Title:         "Charlie X",
        AirDate:       ptrString("1966-09-15"),
        StarDate:      ptrFloat64(1533.6),
    },
    // Add more episodes
}

// Helper functions to create pointers
func ptrString(s string) *string {
    return &s
}

func ptrFloat64(f float64) *float64 {
    return &f
}

seasons/tos/season2.go

Define episodes for Season 2 (if applicable):

go

// seasons/tos/season2.go
package tos

// Define episodes for Season 2
var Season2Episodes = []Episode{
    {
        SeriesID:      1,
        SeasonNumber:  2,
        EpisodeNumber: 1,
        Title:         "Amok Time",
        AirDate:       ptrString("1967-09-15"),
        StarDate:      ptrFloat64(3283.5),
    },
    // Add more episodes
}

2. Set Up Seeding in seed.go

Use the defined series and episodes to seed your database:

go

// seed.go
package main

import (
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/ericthomaca/star-trek-api/seasons/tos" // Adjust import path if necessary
)

func main() {
    // Database connection
    dsn := "user=username password=password dbname=mydb port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Migrate schema
    db.AutoMigrate(&Series{}, &Episode{})

    // Seed series
    for _, s := range seriesList {
        if err := db.Create(&s).Error; err != nil {
            log.Fatalf("Failed to insert series: %v", err)
        }
    }

    // Seed episodes
    allEpisodes := append(tos.Season1Episodes, tos.Season2Episodes...)
    for _, ep := range allEpisodes {
        if err := db.Create(&ep).Error; err != nil {
            log.Printf("Failed to insert episode %v: %v", ep, err)
        }
    }

    log.Println("Data seeding complete.")
}

3. Integration with GitHub

    Add Files to Your Repository: Ensure that series.go, the season files in seasons/tos/, and seed.go are added to your GitHub repository.

    Commit and Push Changes:

    sh

    git add series.go
    git add seasons/tos/season1.go
    git add seasons/tos/season2.go
    git add seed.go
    git commit -m "Add series and episodes data for seeding"
    git push origin main

4. Run the Seeder

To seed your database with the data, run the seed.go file:

sh

go run seed.go

Summary

    Organize Data: Define your series and episodes in separate Go files.
    Create Seeder: Use seed.go to migrate the schema and insert the data into the database.
    Integrate with GitHub: Add, commit, and push your changes to GitHub.
    Run Seeding: Execute the seeding process to populate your database.

By following these steps, you ensure your data is structured and integrated effectively into your Star Trek API project.
