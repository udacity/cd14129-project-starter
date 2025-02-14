package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Vehicle struct {
	ID    uint   `gorm:"primaryKey"`
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
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

	// Auto-migrate the Vehicle model
	err = db.AutoMigrate(&Vehicle{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	// Define a simple GET route
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Database connected!")
	})

	log.Println("Server running on port 8080")
	r.Run(":8080")
}
