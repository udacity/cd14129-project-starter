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

	// TODO: Define the GET /items route

	r.Run(":8080")
}

// itemsHandler handles pagination for items
func itemsHandler(c *gin.Context) {
	// TODO: Implement pagination logic using limit and offset query parameters
    
}
