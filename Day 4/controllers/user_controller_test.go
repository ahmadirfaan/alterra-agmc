package controller

import (
	database_config "alterra-agmc-day4/config/database"
	"alterra-agmc-day4/middleware"
	"alterra-agmc-day4/models/database"
	models "alterra-agmc-day4/models/website"
	"alterra-agmc-day4/repositories"
	"alterra-agmc-day4/services"
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func injectUserController() UserController {
	db := database_config.InitDb()
	repository := repositories.NewUserRepository(db)
	service := services.NewUserService(repository)
	return NewUserController(service)
}

func seedDataUser() {
	db := database_config.InitDb()
	repository := repositories.NewUserRepository(db)
	// inject data
	repository.Save(database.User{
		Id:       GetIntPointer(1),
		Name:     "1",
		Password: "1",
		Email:    "1",
	})
	repository.Save(database.User{
		Id:       GetIntPointer(1),
		Name:     "2",
		Password: "2",
		Email:    "2",
	})
	repository.Save(database.User{
		Id:       GetIntPointer(1),
		Name:     "3",
		Password: "3",
		Email:    "3",
	})
}

func cleanDataUser() {
	db := database_config.InitDb()
	repository := repositories.NewUserRepository(db)
	repository.DeleteUser(1)
	repository.DeleteUser(2)
	repository.DeleteUser(3)
}

func Test_GetUsers_valid(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/v1/users")
	// assertions valid test case
	seedDataUser()
	if assert.NoError(t, injectUserController().GetAllUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
	cleanDataUser()
}

func Test_GetUserById_valid(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	seedDataUser()
	if assert.NoError(t, injectUserController().GetUserById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
	cleanDataUser()
}

func Test_GetUserById_Invalid(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/restricted/v1/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("s")

	if assert.NoError(t, injectUserController().GetUserById(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func Test_Create_User(t *testing.T) {
	userJSON := database.User{
		Name:     "Name",
		Email:    "email",
		Password: "password",
	}

	data, _ := json.Marshal(userJSON)
	reader := bytes.NewReader(data)
	e := echo.New()
	e.Validator = middleware.NewCustomValidator()

	req := httptest.NewRequest(http.MethodPost, "/", reader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/users")

	if assert.NoError(t, injectUserController().CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func Test_Update_User(t *testing.T) {
	userJSON := models.CreateUserRequest{
		Name:     "Name",
		Email:    "email",
		Password: "password",
	}

	data, _ := json.Marshal(userJSON)
	reader := bytes.NewReader(data)
	e := echo.New()

	e.Validator = middleware.NewCustomValidator()

	req := httptest.NewRequest(http.MethodPut, "/", reader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/restricted/v1/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	seedData()
	if assert.NoError(t, injectUserController().UpdateUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func Test_Delete_User(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/restricted/v1/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	seedData()
	if assert.NoError(t, injectUserController().DeleteUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
	cleanData()
}

func Test_LoginUser(t *testing.T) {
	loginJson := models.LoginUserRequest{
		Email:    "email",
		Password: "password",
	}

	data, _ := json.Marshal(loginJson)
	reader := bytes.NewReader(data)
	e := echo.New()
	e.Validator = middleware.NewCustomValidator()

	req := httptest.NewRequest(http.MethodPost, "/", reader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/v1/login")

	if assert.NoError(t, injectUserController().LoginUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
