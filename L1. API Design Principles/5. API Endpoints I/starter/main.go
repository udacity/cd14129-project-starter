package main

import (
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID uint   `json:"user_id"`
}

// Declare a global variable for the GORM database connection.
var db *gorm.DB

func main() {
	var err error

	// Configure a connection to the PostgreSQL database using GORM.
	dsn := "host=localhost user=postgres password=yourpassword dbname=yourdb port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Auto-migrate the Post model to the database.
	db.AutoMigrate(&Post{})

	// TODO: Create a route for creating a single post

	// TODO: Create a route for retrieving a single post

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

// TODO: Implement the handler functions here (i.e., createPost() and getPost())
