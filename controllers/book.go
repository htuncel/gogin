package controllers

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	
	"main/configs"
	"main/models"
)

type BookController struct{}

type BookService interface {
	FindBooks(c *gin.Context)
	FindBook(c *gin.Context)
	CreateBook(c *gin.Context)
	UpdateBook(c *gin.Context)
	DeleteBook(c *gin.Context)
}

func NewBookController() *BookController {
	return new(BookController)
}

// GET /books
// Get all books
func (b *BookController) FindBooks(c *gin.Context) {
	var books []models.Book
	configs.DB.Find(&books)
	
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// GET /books/:id
// Find a book
func (b *BookController) FindBook(c *gin.Context) { // Get model if exist
	var book models.Book
	
	if err := configs.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// POST /books
// Create new book
func (b *BookController) CreateBook(c *gin.Context) {
	// Validate input
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	configs.DB.Create(&book)
	
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// PATCH /books/:id
// Update a book
func (b *BookController) UpdateBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := configs.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	
	// Validate input
	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	configs.DB.Model(&book).Updates(input)
	
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /books/:id
// Delete a book
func (b *BookController) DeleteBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := configs.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	
	configs.DB.Delete(&book)
	
	c.JSON(http.StatusOK, gin.H{"data": true})
}
