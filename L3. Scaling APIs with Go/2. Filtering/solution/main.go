package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"status"`   // e.g., "completed", "pending"
	Priority string `json:"priority"` // e.g., "high", "medium", "low"
	Assignee string `json:"assignee"` // e.g., name of the person assigned
}

var tasks = []Task{
	{ID: 1, Name: "Task 1", Status: "completed", Priority: "high", Assignee: "John"},
	{ID: 2, Name: "Task 2", Status: "pending", Priority: "medium", Assignee: "Alice"},
	{ID: 3, Name: "Task 3", Status: "completed", Priority: "low", Assignee: "Bob"},
	{ID: 4, Name: "Task 4", Status: "pending", Priority: "high", Assignee: "John"},
	{ID: 5, Name: "Task 5", Status: "completed", Priority: "medium", Assignee: "Alice"},
}

func main() {
	r := gin.Default()

	r.GET("/tasks", handleGetTasks)

	log.Println("Server running on port 8080")
	r.Run(":8080")
}

// Handler to filter tasks based on query parameters
func handleGetTasks(c *gin.Context) {
	// Get query parameters from the request
	status := c.Query("status")
	priority := c.Query("priority")
	assignee := c.Query("assignee")

	// Filter tasks based on provided query parameters
	var filteredTasks []Task
	for _, task := range tasks {
		// Check if the task matches all the provided query parameters
		if (status == "" || task.Status == status) &&
			(priority == "" || task.Priority == priority) &&
			(assignee == "" || task.Assignee == assignee) {
			filteredTasks = append(filteredTasks, task)
		}
	}

	// Return the filtered tasks as JSON
	c.JSON(http.StatusOK, filteredTasks)
}
