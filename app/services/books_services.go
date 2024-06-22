package services

import (
	m "books-api/app/models"
	"fmt"
	"gorm.io/gorm"
)

type BooksService struct {
	db *gorm.DB
}

func NewBooksService(db *gorm.DB) *BooksService {
	return &BooksService{db: db}
}

func (bs *BooksService) CreateBook(book *m.Book) (*m.Book, error) {
	if err := bs.db.Create(book).Error; err != nil {
		return nil, fmt.Errorf("failed to create book: %v", err)
	}
	return book, nil
}

func (bs *BooksService) GetAllBooks() ([]m.Book, error) {
	var books []m.Book
	result := bs.db.Find(&books)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve books: %v", result.Error)
	}
	return books, nil
}

func (bs *BooksService) GetBookById(id int64) (*m.Book, error) {
	var book m.Book
	if err := bs.db.Where("id = ?", id).First(&book).Error; err != nil {
		return nil, fmt.Errorf("failed to find book with ID %d: %v", id, err)
	}
	return &book, nil
}

func (bs *BooksService) DeleteBook(id int64) (*m.Book, error) {
	var book m.Book
	if err := bs.db.Where("id = ?", id).First(&book).Error; err != nil {
		return nil, fmt.Errorf("failed to find book with ID %d: %v", id, err)
	}
	if err := bs.db.Delete(&book).Error; err != nil {
		return nil, fmt.Errorf("failed to delete book with ID %d: %v", id, err)
	}
	return &book, nil
}
