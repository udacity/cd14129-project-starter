package main

import (
	"log"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	// TODO: Enable pprof by importing the appropriate package
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

	// TODO: Register pprof routes

	r.GET("/api/data", func(c *gin.Context) {
		simulateWork()
		c.JSON(http.StatusOK, gin.H{"message": "Hello from Gin!"})
	})

	log.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
