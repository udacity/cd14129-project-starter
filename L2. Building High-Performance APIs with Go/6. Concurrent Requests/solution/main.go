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

func handleConcurrentRequests(c *gin.Context) {
	// Create two channels to receive data from the goroutines
	source1Chan := make(chan string)
	source2Chan := make(chan string)

	// Launch two goroutines to simulate concurrent API requests
	go func() {
		time.Sleep(2 * time.Second) // Simulate time taken to fetch data from source 1
		source1Chan <- "Data from Source 1"
	}()

	go func() {
		time.Sleep(3 * time.Second) // Simulate time taken to fetch data from source 2
		source2Chan <- "Data from Source 2"
	}()

	// Wait for both goroutines to finish and collect their results
	source1Data := <-source1Chan
	source2Data := <-source2Chan

	// Combine results and send JSON response
	response := APIResponse{
		Source1: source1Data,
		Source2: source2Data,
	}
	
	c.JSON(http.StatusOK, response)
}
