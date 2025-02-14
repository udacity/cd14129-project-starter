package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// TODO: Apply the customHeaderMiddleware globally
	// Use the Use method to apply the middleware globally

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	r.GET("/goodbye", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Goodbye, World!"})
	})

	r.Run(":8080")
}

// TODO: Implement customHeaderMiddleware function
// customHeaderMiddleware adds a custom header to all responses
func customHeaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Add the custom header

		// TODO: Continue to the next middleware or handler
	}
}
