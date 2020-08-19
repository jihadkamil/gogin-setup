package usecase

import (
	"fmt"
	"gogin/database"
	"gogin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Find Books
func FindBooks(c *gin.Context) {
	var books []models.Book
	database.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// Find A Book
func FindABook(c *gin.Context) {
	var book models.Book

	if err := database.DB.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// Update Book
func UpdateBook(c *gin.Context) {
	var book models.Book

	// get data
	if err := database.DB.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// validate input
	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("errrr 222")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&book).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// Create Book
func CreateBook(c *gin.Context) {
	// validate input
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// create book
	book := models.Book{
		Title:  input.Title,
		Author: input.Author,
	}
	database.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}
