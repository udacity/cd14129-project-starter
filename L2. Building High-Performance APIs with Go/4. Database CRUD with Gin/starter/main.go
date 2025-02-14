package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}

var db *gorm.DB

func main() {
	r := gin.Default()

	// Connect to PostgreSQL using GORM
	dsn := "host=localhost user=postgres password=yourpassword dbname=yourdb port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Auto-migrate the Task model
	err = db.AutoMigrate(&Task{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	// Routes for CRUD operations
	r.POST("/tasks", createTask)
	r.GET("/tasks/:id", getTask)
	r.PUT("/tasks/:id", updateTask)
	r.DELETE("/tasks/:id", deleteTask)

	log.Println("Server running on port 8080")
	r.Run(":8080")
}

// Create a new task
func createTask(c *gin.Context) {
	// TODO: Implement create task functionality
}

// Retrieve a task by its ID
func getTask(c *gin.Context) {
	// TODO: Implement get task functionality
}

// Update a task by its ID
func updateTask(c *gin.Context) {
	// TODO: Implement update task functionality
}

// Delete a task by its ID
func deleteTask(c *gin.Context) {
	// TODO: Implement delete task functionality
}
