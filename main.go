package main

import (
	"groove-galery/controllers"
	"groove-galery/models"

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
	r.Run()
}
