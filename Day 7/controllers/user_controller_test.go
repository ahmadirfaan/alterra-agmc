package controller

import (
	database_config "alterra-agmc-day7/config/database"
	"alterra-agmc-day7/middleware"
	"alterra-agmc-day7/models/database"
	models "alterra-agmc-day7/models/website"
	"alterra-agmc-day7/repositories"
	"alterra-agmc-day7/services"
	"alterra-agmc-day7/utils"
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
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
		Id:        GetIntPointer(1),
		Name:      "1",
		Password:  utils.HashPassword("passwordpassword123"),
		Email:     "email@gmail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	repository.Save(database.User{
		Id:        GetIntPointer(2),
		Name:      "2",
		Password:  "2",
		Email:     "2",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	repository.Save(database.User{
		Id:        GetIntPointer(3),
		Name:      "3",
		Password:  "3",
		Email:     "3",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}

func cleanDataUser() {
	db := database_config.InitDb()
	repository := repositories.NewUserRepository(db)
	repository.DeleteUser(1)
	repository.DeleteUser(2)
	repository.DeleteUser(3)
	repository.DeleteUser(4)
	repository.DeleteUser(5)
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
	token, _, _ := utils.GenerateToken(database.User{
		Id: GetIntPointer(1),
	})
	req.Header.Set("Authorization", "Bearer "+*token)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/restricted/v1/users/:id")
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
	seedDataUser()
	if assert.NoError(t, injectUserController().GetUserById(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
	cleanDataUser()
}

func Test_Create_User(t *testing.T) {
	atoi := strconv.Itoa(rand.Int())
	userJSON := models.CreateUserRequest{
		Name:     "Name",
		Email:    "email" + atoi + "@yahoo.com",
		Password: "password123",
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
		Email:    "email123@gmail.com",
		Password: "passwordpassword123",
	}

	data, _ := json.Marshal(userJSON)
	reader := bytes.NewReader(data)
	e := echo.New()

	e.Validator = middleware.NewCustomValidator()

	req := httptest.NewRequest(http.MethodPut, "/", reader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	token, _, _ := utils.GenerateToken(database.User{
		Id: GetIntPointer(1),
	})
	req.Header.Set("Authorization", "Bearer "+*token)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/restricted/v1/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	seedDataUser()
	if assert.NoError(t, injectUserController().UpdateUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
	cleanDataUser()
}

func Test_Delete_User(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	token, _, _ := utils.GenerateToken(database.User{
		Id: GetIntPointer(1),
	})
	req.Header.Set("Authorization", "Bearer "+*token)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/restricted/v1/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	seedDataUser()
	if assert.NoError(t, injectUserController().DeleteUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
	cleanDataUser()
}

func Test_LoginUser(t *testing.T) {
	loginJson := models.LoginUserRequest{
		Email:    "email@gmail.com",
		Password: "passwordpassword123",
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
	seedDataUser()
	if assert.NoError(t, injectUserController().LoginUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
	cleanDataUser()
}
