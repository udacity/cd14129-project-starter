package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func BenchmarkAPIData(b *testing.B) {
	// Use the setupRouter function from main.go to ensure consistent routes
	router := setupRouter()

	// Perform the benchmark
	b.ResetTimer() // Reset timer to exclude setup overhead
	for i := 0; i < b.N; i++ {
		// TODO: Create a new HTTP request for the /api/data route

		// TODO: Initializes a response recorder 

		// TODO: Serve the request using the Gin router

		// Check the response status
		if resp.Code != http.StatusOK {
			b.Errorf("Expected status %v, got %v", http.StatusOK, resp.Code)
		}
	}
}
