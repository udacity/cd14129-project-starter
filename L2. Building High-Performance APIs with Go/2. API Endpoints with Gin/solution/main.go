package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
}

func main() {
	router := gin.Default()

	// Define the GET /users route and handler
	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users) // Return the list of users in JSON format
	})

	// Define the POST /users route and handler
	router.POST("/users", func(c *gin.Context) {
		var newUser User

		// Bind the incoming JSON to newUser struct
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Assign a new ID to the user and append to the users slice
		newUser.ID = len(users) + 1
		users = append(users, newUser)

		// Return the newly created user with status 201
		c.JSON(http.StatusCreated, newUser)
	})

	log.Println("Server running on port 8080")
	router.Run(":8080")
}
