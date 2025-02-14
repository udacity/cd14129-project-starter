package main

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin" // Import Gin package
)

func main() {
	// Initialize the Gin router
	router := gin.Default()

	// Define a simple GET route for "/"
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Gin!") // Respond with a simple message
	})

	log.Println("Server running on port 8080")

	// Start the Gin server on port 8080
	router.Run(":8080")
}
