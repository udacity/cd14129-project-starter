package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var items = []string{
	"Item 1", "Item 2", "Item 3", "Item 4", "Item 5",
	"Item 6", "Item 7", "Item 8", "Item 9", "Item 10",
	"Item 11", "Item 12", "Item 13", "Item 14", "Item 15",
}

func main() {
	r := gin.Default()

    // Define the GET /items route
	r.GET("/items", itemsHandler)

	r.Run(":8080")
}

func itemsHandler(c *gin.Context) {
	// Retrieve limit and offset from query parameters
	limitParam := c.Query("limit")
	offsetParam := c.Query("offset")

	limit := 5 // Default limit
	offset := 0 // Default offset

	// Parse limit and offset parameters
	if limitParam != "" {
		var err error
		limit, err = strconv.Atoi(limitParam)
		if err != nil || limit <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
			return
		}
	}

	if offsetParam != "" {
		var err error
		offset, err = strconv.Atoi(offsetParam)
		if err != nil || offset < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
			return
		}
	}

	if offset >= len(items) {
		c.JSON(http.StatusOK, gin.H{"items": []string{}})
		return
	}

	end := offset + limit
	if end > len(items) {
		end = len(items)
	}

	// Return the paginated items
	c.JSON(http.StatusOK, gin.H{"items": items[offset:end]})
}
