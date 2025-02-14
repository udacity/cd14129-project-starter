package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Apply the customHeaderMiddleware globally
	// Use the Use method to apply the middleware globally
	r.Use(customHeaderMiddleware())

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	r.GET("/goodbye", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Goodbye, World!"})
	})

	r.Run(":8080")
}

// customHeaderMiddleware adds a custom header to all responses
func customHeaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add the custom header
		c.Header("X-Custom-Header", "Middleware-Active")

		// Continue to the next middleware or handler
		c.Next()
	}
}

/* 
Below is an example of a GET request to the /hello endpoint and 
what the response would look like, including the headers. 

Note that in the response, the middleware successfully added the
X-Custom-Header: Middleware-Active header, as expected.

Request: 
curl -i http://localhost:8080/hello

Response:
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
X-Custom-Header: Middleware-Active
Date: Tue, 15 Oct 2024 01:02:03 GMT
Content-Length: 27

{
  "message": "Hello, World!"
}
*/
