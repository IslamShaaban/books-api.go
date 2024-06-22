package routes

import (
	"books-api/routes/books"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) *gin.Engine {
	r.Use(gin.Recovery())
	books.BooksRoutes(r, db)
	return r
}
