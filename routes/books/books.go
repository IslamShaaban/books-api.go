package books

import (
	"books-api/app/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BooksRoutes(r *gin.Engine, db *gorm.DB) *gin.RouterGroup {
	booksController := controllers.NewBooksController(db)
	booksGroup := r.Group("/books")
	{
		booksGroup.POST("", booksController.Create)
		booksGroup.GET("", booksController.Index)
		booksGroup.GET("/:id", booksController.Show)
		booksGroup.DELETE("/:id", booksController.Delete)
	}
	return booksGroup
}
