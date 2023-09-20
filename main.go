package main

import (
	"groove-gallery/controllers"
	"groove-gallery/models"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/albums", controllers.FindAlbums)
	r.GET("/albums/:id", controllers.FindAlbum)
	r.POST("/albums", controllers.CreateAlbum)
	r.PATCH("/albums/:id", controllers.UpdateAlbum)
	r.DELETE("/albums/:id", controllers.DeleteAlbum)

	// Run the server
	groovePort := os.Getenv("GROOVE_PORT")

	if len(groovePort) == 0 {
		groovePort = "8080"
	}
	r.Run(os.Getenv("GROOVE_PORT"))
}
