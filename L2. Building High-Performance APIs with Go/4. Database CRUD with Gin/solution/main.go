package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Task represents a simple task model in the database
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
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := db.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create task"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// Retrieve a task by its ID
func getTask(c *gin.Context) {
	id := c.Param("id")
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// Update a task by its ID
func updateTask(c *gin.Context) {
	id := c.Param("id")
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var updatedTask Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	task.Title = updatedTask.Title
	task.Description = updatedTask.Description
	task.IsCompleted = updatedTask.IsCompleted

	if err := db.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// Delete a task by its ID
func deleteTask(c *gin.Context) {
	id := c.Param("id")
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err := db.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
