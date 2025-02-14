package main

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string
	Description string
}

// Declare the database connection
var db *gorm.DB

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Task Manager API!")
	})

	// Set up a connection to the PostgreSQL database using GORM
	dsn := "host=localhost user=postgres password=yourpassword dbname=yourdb port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	// Auto-migrate the Task model to the database
	db.AutoMigrate(&Task{})

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
