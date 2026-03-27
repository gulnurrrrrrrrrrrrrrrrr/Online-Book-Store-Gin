package handlers

import (
	"Online-Book-Store-Gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var books = make(map[int]models.Book)
var nextBookID = 1

func GetBooks(c *gin.Context) {
	var bookList []models.Book
	for _, book := range books {
		bookList = append(bookList, book)
	}
	c.JSON(http.StatusOK, bookList)
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if book.Title == "" || book.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required and Price must be > 0"})
		return
	}

	book.ID = nextBookID
	nextBookID++
	books[book.ID] = book

	c.JSON(http.StatusCreated, book)
}

func GetBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	book, exists := books[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if _, exists := books[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	var updatedBook models.Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updatedBook.Title == "" || updatedBook.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required and Price must be > 0"})
		return
	}

	updatedBook.ID = id
	books[id] = updatedBook

	c.JSON(http.StatusOK, updatedBook)
}

func DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if _, exists := books[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	delete(books, id)
	c.JSON(http.StatusNoContent, nil)
}
