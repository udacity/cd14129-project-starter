/*

Key optimization techniques in this solution:

- Reducing memory allocations: By avoiding redundant memory allocations, we reduce
 the amount of memory used per operation (B/op).

- Improving algorithmic efficiency: Recomputing the same value multiple times
slows down the execution. By computing it once and reusing it, we reduce the
computational complexity.

*/

package main

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Optimized API endpoint
	r.GET("/compute", optimizedCompute)

	r.Run(":8080")
}

// Optimized compute function
func optimizedCompute(c *gin.Context) {
	sum := 0.0

	for i := 0; i < 1000000; i++ {
		// Avoid computing sqrt(i) twice
		sqrt := math.Sqrt(float64(i)) // Compute sqrt(i) once
		sum += sqrt * sqrt            // Reuse computed sqrt(i)
	}

	c.JSON(http.StatusOK, gin.H{"result": sum})
}
