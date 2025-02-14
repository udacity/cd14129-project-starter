package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	"net/http"
	"time"
	_ "Photos/docs"
)

type Photo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

// ErrorResponse is used for returning error responses
type ErrorResponse struct {
	Error string `json:"error"`
}

var photos = []Photo{
	{ID: "1", Title: "Sunset", URL: "https://example.com/sunset.jpg"},
	{ID: "2", Title: "Mountain", URL: "https://example.com/mountain.jpg"},
	{ID: "3", Title: "Ocean", URL: "https://example.com/ocean.jpg"},
}

func getAllPhotos(c *gin.Context) {
	c.JSON(http.StatusOK, photos)
}

func createPhoto(c *gin.Context) {
	var newPhoto Photo
	if err := c.ShouldBindJSON(&newPhoto); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	newPhoto.ID = time.Now().Format("20060102150405") // Simple ID based on timestamp
	photos = append(photos, newPhoto)

	c.JSON(http.StatusCreated, newPhoto)
}

func main() {
	r := gin.Default()

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Modify routes for photos
	r.GET("/photos", getAllPhotos)
	r.POST("/photo", createPhoto)

	r.Run(":8080")
}
