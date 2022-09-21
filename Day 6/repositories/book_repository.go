package repositories

import (
	"alterra-agmc-day6/models/database"
	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(book *database.Book) error
	GetBookById(id int) (database.Book, error)
	UpdateBook(book *database.Book, id int) error
	DeleteBook(id int) error
	GetAllBooks(offset int) ([]database.Book, error)
}

type bookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		DB: db,
	}
}

func (b bookRepository) CreateBook(book *database.Book) error {
	if err := b.DB.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (b bookRepository) UpdateBook(book *database.Book, id int) error {
	if err := b.DB.Model(book).Where("id = ?", id).Updates(book).Error; err != nil {
		return err
	}
	return nil
}

func (b bookRepository) DeleteBook(id int) error {
	if err := b.DB.Delete(&database.Book{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (b bookRepository) GetBookById(id int) (database.Book, error) {
	var book database.Book
	result := b.DB.Where("id = ?", id).First(&book)
	return book, result.Error
}

func (b bookRepository) GetAllBooks(offset int) ([]database.Book, error) {
	var books []database.Book
	result := b.DB.Limit(25).Offset(offset).Find(&books)
	return books, result.Error
}
