package main

import (
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Inefficient API endpoint
	r.GET("/compute", inefficientCompute)

	r.Run(":8080")
}

// Inefficient compute function
func inefficientCompute(c *gin.Context) {
	sum := 0.0

	// Unnecessarily recomputing values multiple times
	for i := 0; i < 1000000; i++ {
		// TODO: Optimize the code within this "for" loop
		// Hint: You may want to avoid computing sqrt(i) twice
		sum += math.Sqrt(float64(i)) * math.Sqrt(float64(i)) // Unnecessary double calculation
		time.Sleep(1 * time.Nanosecond)                      // Simulated delay
	}

	c.JSON(http.StatusOK, gin.H{"result": sum})
}
