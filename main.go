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
	var groovePort string

	if len(os.Getenv("GROOVE_PORT")) == 0 {
		groovePort = ":8080"
	} else {
		groovePort = ":" + os.Getenv("GROOVE_PORT")
	}

	err := r.Run(groovePort)
	if err != nil {
		panic("[ERROR] failed  to start Gin server due to " + err.Error())
	}
}
