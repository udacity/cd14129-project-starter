package main

import (
	"log"
	"net/http"
	_ "net/http/pprof" // pprof enabled
	"time"
	"github.com/gin-gonic/gin"
)

func simulateWork() {
	// Simulate CPU-intensive work
	for i := 0; i < 10000000; i++ {
		_ = i * i
	}

	// Simulate memory-intensive work by allocating a large slice
	data := make([]byte, 10<<20) // 10MB
	for i := range data {
		data[i] = byte(i % 256)
	}
	_ = data // Prevent compiler optimization
}

func main() {
	r := gin.Default()

	// Register pprof routes
	go func() {
		log.Println("pprof enabled on port 6060")
		http.ListenAndServe(":6060", nil) // Serve pprof routes on a separate port
	}()

	r.GET("/api/data", func(c *gin.Context) {
		simulateWork()
		c.JSON(http.StatusOK, gin.H{"message": "Hello from Gin!"})
	})

	log.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
