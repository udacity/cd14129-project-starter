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
		c.JSON(http.StatusOK, todos)
	})

	// Endpoint to retrieve a single todo by ID with caching
	router.GET("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Check the cache for the todo
		cacheMutex.RLock() // Read lock for cache
		if todo, found := cache[id]; found {
			cacheMutex.RUnlock()
			c.JSON(http.StatusOK, todo)
			return
		}
		cacheMutex.RUnlock()

		// If not found in cache, search the todos list
		for _, todo := range todos {
			if todo.ID == id {
				// Add the todo to the cache for future requests
				cacheMutex.Lock() // Write lock for cache
				cache[id] = todo
				cacheMutex.Unlock()

				c.JSON(http.StatusOK, todo)
				return
			}
		}

		// If the todo was not found, return a 404 error
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
	})

	router.Run(":8080")
}
