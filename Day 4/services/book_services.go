package services

import (
	"alterra-agmc-day4/models/database"
	models "alterra-agmc-day4/models/website"
	"alterra-agmc-day4/repositories"
	"time"
)

type BookService interface {
	CreateNewBook(request models.CreateBookRequest) error
	GetBookById(id int) (database.Book, error)
	UpdateBook(book *models.CreateBookRequest, id int) error
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

func (b *bookService) CreateNewBook(request models.CreateBookRequest) error {
	newBook := database.Book{
		Title:     request.Title,
		ISBN:      request.ISBN,
		Writer:    request.Writer,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := b.bookRepository.CreateBook(&newBook)
	return err
}

func (b *bookService) GetBookById(id int) (database.Book, error) {
	book, err := b.bookRepository.GetBookById(id)
	return book, err
}
func (b *bookService) UpdateBook(book *models.CreateBookRequest, id int) error {
	newBook := &database.Book{
		Title:     book.Title,
		ISBN:      book.ISBN,
		Writer:    book.Writer,
		UpdatedAt: time.Now(),
	}
	err := b.bookRepository.UpdateBook(newBook, id)
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
