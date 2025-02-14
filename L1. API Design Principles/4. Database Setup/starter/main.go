package main

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Task struct is our database model
type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string
	Description string
}

func main() {
	// TODO: Declare the database connection

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Task Manager API!")
	})

	// TODO: Set up a connection to the PostgreSQL database using GORM

	// TODO: Auto-migrate the Task model to the database

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
