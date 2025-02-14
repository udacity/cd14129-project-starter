package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"sync"
)

// Todo represents a simple todo with an ID and a title
type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var todos = []Todo{
	{ID: "1", Title: "Buy groceries"},
	{ID: "2", Title: "Walk the dog"},
	{ID: "3", Title: "Do the laundry"},
}

// In-memory cache (using a map)
var cache = make(map[string]Todo)
var cacheMutex sync.RWMutex // Mutex for thread safety when accessing cache

func main() {
	router := gin.Default()

	// Endpoint to retrieve all todos
	router.GET("/todos", func(c *gin.Context) {
		// TODO: Return all todos
	})

	// Endpoint to retrieve a single todo by ID
	router.GET("/todos/:id", func(c *gin.Context) {
		// TODO: Retrieve todo from cache first, then fallback to searching the list
	})
	
	router.Run(":8080")
}
