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

	// TODO: Define the GET /users route and handler

	// TODO: Define the POST /users route and handler

	log.Println("Server running on port 8080")
	router.Run(":8080")
}
