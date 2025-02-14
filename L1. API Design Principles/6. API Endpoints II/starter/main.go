package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

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
	http.HandleFunc("/posts", createPost)
	// TODO: Create a route for retrieving a single post
	http.HandleFunc("/posts/", getPost)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

// Handler function to create a new post (POST /posts)
func createPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var post Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Insert the post into the database
	if err := db.Create(&post).Error; err != nil {
		http.Error(w, "Could not create post", http.StatusInternalServerError)
		return
	}

	// Return the newly created post as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// Handler function to get a post by ID (GET /posts/{id})
func getPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract the ID from the URL path
	idStr := strings.TrimPrefix(r.URL.Path, "/posts/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	var post Post
	// Retrieve the post from the database
	if err := db.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Post not found", http.StatusNotFound)
		} else {
			http.Error(w, "Could not retrieve post", http.StatusInternalServerError)
		}
		return
	}

	// Return the post as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}
