package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// setupRouter initializes the Gin router and defines all routes.
// This function is reused in tests for consistency.
func setupRouter() *gin.Engine {
	r := gin.Default()

	// A sample API endpoint that simulates some processing time
	r.GET("/api/data", func(c *gin.Context) {
		time.Sleep(100 * time.Millisecond) // Simulate a delay
		c.JSON(http.StatusOK, gin.H{"message": "Hello from Gin!"})
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
