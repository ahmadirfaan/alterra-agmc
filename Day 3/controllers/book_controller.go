package controller

import (
	"alterra-agmc-dynamic-crud/models/database"
	models "alterra-agmc-dynamic-crud/models/website"
	"alterra-agmc-dynamic-crud/services"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type BookController interface {
	GetAllBooks(c echo.Context) error
	CreateBook(c echo.Context) error
	UpdateBook(c echo.Context) error
	DeleteBook(c echo.Context) error
	GetBookById(c echo.Context) error
}

type bookController struct {
	BookService services.BookService
}

func NewBookController(s services.BookService) BookController {
	return bookController{
		BookService: s,
	}
}

func (uc bookController) GetAllBooks(c echo.Context) error {
	queryParam := c.QueryParam("page")
	page, _ := strconv.Atoi(queryParam)
	books, err := uc.BookService.GetAllBooks(page)
	if err != nil {
		return wrapperResponse(http.StatusInternalServerError, "Error", nil).ConvertDataJSON(c.Response())
	}
	return wrapperResponse(http.StatusOK, "Success get all books", books).ConvertDataJSON(c.Response())

}

func (uc bookController) CreateBook(c echo.Context) error {

	book := models.CreateBookRequest{}
	err := json.NewDecoder(c.Request().Body).Decode(&book)
	if err != nil {
		return wrapperResponse(http.StatusBadRequest, "Error", nil).ConvertDataJSON(c.Response())
	}

	newBook := database.Book{
		Title:  book.Title,
		ISBN:   book.ISBN,
		Writer: book.Writer,
	}

	errServices := uc.BookService.CreateNewBook(newBook)
	if errServices != nil {
		return wrapperResponse(http.StatusInternalServerError, "Error", nil).ConvertDataJSON(c.Response())
	}
	return wrapperResponse(http.StatusOK, "Success create book", newBook).ConvertDataJSON(c.Response())
}

func (uc bookController) UpdateBook(c echo.Context) error {
	book := models.CreateBookRequest{}
	err := json.NewDecoder(c.Request().Body).Decode(&book)
	id, _ := strconv.Atoi(c.Param("id"))
	if err != nil {
		return wrapperResponse(http.StatusBadRequest, "Error", nil).ConvertDataJSON(c.Response())
	}
	newBook := &database.Book{
		Title:  book.Title,
		ISBN:   book.ISBN,
		Writer: book.Writer,
	}
	errServices := uc.BookService.UpdateBook(newBook, id)
	if errServices != nil {
		return wrapperResponse(http.StatusInternalServerError, "Error", nil).ConvertDataJSON(c.Response())
	}
	return wrapperResponse(http.StatusOK, "Success update book", newBook).ConvertDataJSON(c.Response())
}

func (uc bookController) DeleteBook(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return wrapperResponse(http.StatusBadRequest, "Error", nil).ConvertDataJSON(c.Response())
	}

	errServices := uc.BookService.DeleteBook(id)
	if errServices != nil {
		return wrapperResponse(http.StatusInternalServerError, "Error", nil).ConvertDataJSON(c.Response())
	}
	return wrapperResponse(http.StatusOK, "Success Delete book", nil).ConvertDataJSON(c.Response())
}

func (uc bookController) GetBookById(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return wrapperResponse(http.StatusBadRequest, "Error", nil).ConvertDataJSON(c.Response())
	}

	book, errServices := uc.BookService.GetBookById(id)
	if errServices != nil {
		return wrapperResponse(http.StatusInternalServerError, "Error", nil).ConvertDataJSON(c.Response())
	}
	return wrapperResponse(http.StatusOK, "Success Get Book", book).ConvertDataJSON(c.Response())
}
