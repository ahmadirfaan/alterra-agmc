package services

import (
	"alterra-agmc-dynamic-crud/models/database"
	"alterra-agmc-dynamic-crud/repositories"
)

type BookService interface {
	CreateNewBook(request database.Book) error
	GetBookById(id int) (database.Book, error)
	UpdateBook(book *database.Book, id int) error
	DeleteBook(id int) error
	GetAllBooks(page int) ([]database.Book, error)
}

type bookService struct {
	bookRepository repositories.BookRepository
}

func NewBookService(br repositories.BookRepository) BookService {
	return &bookService{
		bookRepository: br,
	}
}

func (b *bookService) CreateNewBook(request database.Book) error {
	err := b.bookRepository.CreateBook(&request)
	return err
}

func (b *bookService) GetBookById(id int) (database.Book, error) {
	book, err := b.bookRepository.GetBookById(id)
	return book, err
}
func (b *bookService) UpdateBook(book *database.Book, id int) error {
	err := b.bookRepository.UpdateBook(book, id)
	return err
}

func (b *bookService) DeleteBook(id int) error {
	err := b.bookRepository.DeleteBook(id)
	return err
}

func (b *bookService) GetAllBooks(page int) ([]database.Book, error) {
	var offset = 0
	if page > 1 {
		offset = 25 * page
	}
	books, err := b.bookRepository.GetAllBooks(offset)
	return books, err
}
