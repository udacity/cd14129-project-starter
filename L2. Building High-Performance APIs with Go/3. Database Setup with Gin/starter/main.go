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

	// TODO: Connect to PostgreSQL using GORM

	// TODO: Auto-migrate the Vehicle model

	// TODO: Define a simple GET route

	log.Println("Server running on port 8080")
	r.Run(":8080")
}
