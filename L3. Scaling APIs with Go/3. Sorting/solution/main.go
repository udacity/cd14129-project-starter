package main

import (
	"log"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Director    string  `json:"director"`
	ReleaseYear int     `json:"releaseYear"`
	Rating      float32 `json:"rating"` // Rating out of 10
}

var movies = []Movie{
	{ID: 1, Title: "Inception", Director: "Christopher Nolan", ReleaseYear: 2010, Rating: 8.8},
	{ID: 2, Title: "The Matrix", Director: "Lana Wachowski", ReleaseYear: 1999, Rating: 8.7},
	{ID: 3, Title: "Interstellar", Director: "Christopher Nolan", ReleaseYear: 2014, Rating: 8.6},
	{ID: 4, Title: "The Godfather", Director: "Francis Ford Coppola", ReleaseYear: 1972, Rating: 9.2},
	{ID: 5, Title: "Parasite", Director: "Bong Joon-ho", ReleaseYear: 2019, Rating: 8.6},
}

func main() {
	r := gin.Default()

	r.GET("/movies", handleGetMovies)

	log.Println("Server running on port 8080")
	r.Run(":8080")
}

// Handler to sort movies based on query parameter "sortBy"
func handleGetMovies(c *gin.Context) {
	// Get query parameters from the request
	sortBy := c.Query("sortBy")

	// Sort movies based on the "sortBy" parameter
	switch sortBy {
	case "title":
		sort.Slice(movies, func(i, j int) bool { return movies[i].Title < movies[j].Title })
	case "director":
		sort.Slice(movies, func(i, j int) bool { return movies[i].Director < movies[j].Director })
	case "releaseYear":
		sort.Slice(movies, func(i, j int) bool { return movies[i].ReleaseYear < movies[j].ReleaseYear })
	case "rating":
		sort.Slice(movies, func(i, j int) bool { return movies[i].Rating > movies[j].Rating }) // Higher rating first
	default:
		sort.Slice(movies, func(i, j int) bool { return movies[i].ID < movies[j].ID }) // Default sort by ID
	}

	// Return the sorted movies as JSON
	c.JSON(http.StatusOK, movies)
}
