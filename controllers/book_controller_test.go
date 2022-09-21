package controller

import (
	database_config "alterra-agmc-day7/config/database"
	"alterra-agmc-day7/middleware"
	"alterra-agmc-day7/models/database"
	models "alterra-agmc-day7/models/website"
	"alterra-agmc-day7/repositories"
	"alterra-agmc-day7/services"
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func injectGetBookController() BookController {
	db := database_config.InitDb()
	repository := repositories.NewBookRepository(db)
	service := services.NewBookService(repository)
	return NewBookController(service)
}

func seedData() {
	db := database_config.InitDb()
	repository := repositories.NewBookRepository(db)
	// inject data
	repository.CreateBook(&database.Book{
		Id:        GetIntPointer(1),
		Title:     "1",
		Writer:    "1",
		ISBN:      "1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	repository.CreateBook(&database.Book{
		Id:     GetIntPointer(2),
		Title:  "2",
		Writer: "2",
		ISBN:   "2",
	})
	repository.CreateBook(&database.Book{
		Id:        GetIntPointer(3),
		Title:     "3",
		Writer:    "3",
		ISBN:      "3",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}

func cleanData() {
	db := database_config.InitDb()
	repository := repositories.NewBookRepository(db)
	repository.DeleteBook(1)
	repository.DeleteBook(2)
	repository.DeleteBook(3)
}

func GetIntPointer(value int) *int {
	return &value
}

func Test_GetBooks_valid(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/v1/books")
	// assertions valid test case
	if assert.NoError(t, injectGetBookController().GetAllBooks(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func Test_GetBookById_valid(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	seedData()
	if assert.NoError(t, injectGetBookController().GetBookById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
	cleanData()
}

func Test_GetBookById_invalid(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("s")

	if assert.NoError(t, injectGetBookController().GetBookById(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func Test_Create_Book(t *testing.T) {
	bookJSON := models.CreateBookRequest{
		Title:  "title",
		Writer: "writer",
		ISBN:   "123123-84848-8484",
	}

	data, _ := json.Marshal(bookJSON)
	reader := bytes.NewReader(data)
	e := echo.New()
	e.Validator = middleware.NewCustomValidator()

	req := httptest.NewRequest(http.MethodPost, "/", reader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/restricted/v1/books")

	if assert.NoError(t, injectGetBookController().CreateBook(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func Test_Update_Book(t *testing.T) {
	bookJSON := models.CreateBookRequest{
		Title:  "title update",
		Writer: "author update",
		ISBN:   "123123-84848-8484-update",
	}

	data, _ := json.Marshal(bookJSON)
	reader := bytes.NewReader(data)
	e := echo.New()

	e.Validator = middleware.NewCustomValidator()

	req := httptest.NewRequest(http.MethodPut, "/", reader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/restricted/v1/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	seedData()
	if assert.NoError(t, injectGetBookController().UpdateBook(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func Test_Delete_Book(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/restricted/v1/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	seedData()
	if assert.NoError(t, injectGetBookController().DeleteBook(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
	cleanData()
}
