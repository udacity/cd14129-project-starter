package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type APIResponse struct {
	Source1 string `json:"source1"`
	Source2 string `json:"source2"`
}

func main() {
	r := gin.Default()

	r.GET("/data", handleConcurrentRequests)

	r.Run(":8080")
}

// handleConcurrentRequests handles concurrent API requests
func handleConcurrentRequests(c *gin.Context) {
	// TODO: Create two channels to receive data from the goroutines

	// TODO: Launch two goroutines to simulate concurrent API requests

	// TODO: Wait for both goroutines to finish and collect their results

	// TODO: Combine results and send JSON response

	// Uncomment the next line once you complete the above steps
	// c.JSON(http.StatusOK, response)
}
