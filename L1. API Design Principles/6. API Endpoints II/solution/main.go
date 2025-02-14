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

	http.HandleFunc("/posts", createPost)      // For creating a new post (POST)
	http.HandleFunc("/posts", getAllPosts)     // For retrieving all posts (GET)
	http.HandleFunc("/posts/", getPostByID)    // For retrieving a single post (GET)
	http.HandleFunc("/posts/", updatePostByID) // For updating a post (PATCH)
	http.HandleFunc("/posts/", deletePostByID) // For deleting a post (DELETE)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

// Handler to create a new post (POST /posts)
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// Handler to retrieve all posts (GET /posts)
func getAllPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var posts []Post
	if err := db.Find(&posts).Error; err != nil {
		http.Error(w, "Could not retrieve posts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// Handler to retrieve a post by ID (GET /posts/{id})
func getPostByID(w http.ResponseWriter, r *http.Request) {
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
	if err := db.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Post not found", http.StatusNotFound)
		} else {
			http.Error(w, "Could not retrieve post", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// Handler to update a post by ID (PATCH /posts/{id})
func updatePostByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
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

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Update the post in the database
	if err := db.Model(&Post{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		http.Error(w, "Could not update post", http.StatusInternalServerError)
		return
	}

	// Return 204 No Content on successful update
	w.WriteHeader(http.StatusNoContent)
}

// Handler to delete a post by ID (DELETE /posts/{id})
func deletePostByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
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

	if err := db.Delete(&Post{}, id).Error; err != nil {
		http.Error(w, "Could not delete post", http.StatusInternalServerError)
		return
	}

	// Return 204 No Content on successful deletion
	w.WriteHeader(http.StatusNoContent)
}
