package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"static-crud/model"
	"strconv"
	"time"
)

func GetAllBooks(c echo.Context) error {
	return wrapperResponse(http.StatusOK, "success get all books", model.BookData).ConvertDataJSON(c.Response())
}

func GetBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, book := range model.BookData {
		if book.ID == uint(id) {
			return wrapperResponse(http.StatusOK, "success get detail book", book).ConvertDataJSON(c.Response())
		}
	}
	return wrapperResponse(http.StatusBadRequest, "failed get detail", &model.Book{}).ConvertDataJSON(c.Response())
}

func CreateNewBook(c echo.Context) error {
	book := model.Book{}
	err := json.NewDecoder(c.Request().Body).Decode(&book)
	if err != nil {
		return wrapperResponse(http.StatusBadRequest, "failed create book", &model.Book{}).ConvertDataJSON(c.Response())
	}
	book.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	model.BookData = append(model.BookData, book)

	return wrapperResponse(http.StatusCreated, "success create book", &model.Book{}).ConvertDataJSON(c.Response())
}

func UpdateBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	requestBook := model.Book{}
	err := json.NewDecoder(c.Request().Body).Decode(&requestBook)
	if err != nil {
		return wrapperResponse(http.StatusBadRequest, "failed update book", &model.Book{}).ConvertDataJSON(c.Response())
	}
	for i, book := range model.BookData {
		if book.ID == uint(id) {
			requestBook.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			model.BookData[i] = requestBook
			return wrapperResponse(http.StatusOK, "success update book", &model.Book{}).ConvertDataJSON(c.Response())
		}
	}

	return wrapperResponse(http.StatusBadRequest, "failed update book", &model.Book{}).ConvertDataJSON(c.Response())
}

func DeleteBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, book := range model.BookData {
		if book.ID == uint(id) {
			model.BookData[i] = model.BookData[0]
			model.BookData = model.BookData[1:]
			return wrapperResponse(http.StatusOK, "success delete book", &model.Book{}).ConvertDataJSON(c.Response())
		}
	}

	return wrapperResponse(http.StatusBadRequest, "failed delete book", &model.Book{}).ConvertDataJSON(c.Response())
}

func wrapperResponse(code int, message string, response interface{}) *model.HTTPResponse {
	newResponse := new(model.HTTPResponse)
	newResponse.Code = code
	newResponse.Message = message
	newResponse.Data = response
	return newResponse
}
