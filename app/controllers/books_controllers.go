package controllers

import (
	"books-api/app/models"
	"books-api/app/services"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type BooksController struct {
	bookService *services.BooksService
}

func NewBooksController(db *gorm.DB) *BooksController {
	return &BooksController{
		bookService: services.NewBooksService(db),
	}
}

func (bc *BooksController) Create(c *gin.Context) {
	var book models.Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdBook, err := bc.bookService.CreateBook(&book)
	if err == nil {
		c.JSON(http.StatusCreated, createdBook)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
}

func (bc *BooksController) Index(c *gin.Context) {
	books, err := bc.bookService.GetAllBooks()
	if err == nil {
		c.JSON(http.StatusOK, books)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
}

func (bc *BooksController) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := bc.bookService.GetBookById(int64(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to fetch book: %v", err)})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (bc *BooksController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	deletedBook, err := bc.bookService.DeleteBook(int64(id))
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Book deleted", "book": deletedBook})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
}
