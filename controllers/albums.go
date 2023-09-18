package controllers

import (
	"fmt"
	"groove-gallery/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateAlbumInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Genre  string `json:"genre" binding:"required"`
}

type UpdateAlbumInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre" binding:"required"`
}

// GET /books
// Find all books
func FindAlbums(c *gin.Context) {
	var albums []models.Album
	models.DB.Find(&albums)

	c.JSON(http.StatusOK, gin.H{"data": albums})
}

// GET /books/:id
// Find a book
func FindAlbum(c *gin.Context) {
	// Get model if exist
	var album models.Album
	if err := models.DB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": album})
}

// POST /books
// Create new book
func CreateAlbum(c *gin.Context) {
	fmt.Println("Pruebaaaaaa")
	// Validate input
	var input CreateAlbumInput
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("AAAAAAAAAAAAAAAAA")
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	album := models.Album{Title: input.Title, Author: input.Author, Genre: input.Genre}
	models.DB.Create(&album)

	c.JSON(http.StatusOK, gin.H{"data": album})
}

// PATCH /books/:id
// Update a book
func UpdateAlbum(c *gin.Context) {
	// Get model if exist
	var album models.Album
	if err := models.DB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateAlbumInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&album).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": album})
}

// DELETE /books/:id
// Delete a book
func DeleteAlbum(c *gin.Context) {
	// Get model if exist
	var album models.Album
	if err := models.DB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&album)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
