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

type UserController interface {
	GetAllUsers(c echo.Context) error
	CreateUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
	GetUserById(c echo.Context) error
}

type userController struct {
	UserService services.UserService
}

func NewUserController(s services.UserService) UserController {
	return userController{
		UserService: s,
	}
}

func (uc userController) GetAllUsers(c echo.Context) error {
	queryParam := c.QueryParam("page")
	page, _ := strconv.Atoi(queryParam)
	users, err := uc.UserService.GetAllUsers(page)
	if err != nil {
		return wrapperResponse(http.StatusInternalServerError, "Error", nil).ConvertDataJSON(c.Response())
	}
	return wrapperResponse(http.StatusOK, "Success get all users", users).ConvertDataJSON(c.Response())

}

func (uc userController) CreateUser(c echo.Context) error {

	user := models.CreateUserRequest{}
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		return wrapperResponse(http.StatusBadRequest, "Error", nil).ConvertDataJSON(c.Response())
	}

	newUser := database.User{
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
	}
	errServices := uc.UserService.CreateNewUser(newUser)
	if errServices != nil {
		return wrapperResponse(http.StatusInternalServerError, "Error", nil).ConvertDataJSON(c.Response())
	}
	return wrapperResponse(http.StatusOK, "Success create user", user.Email).ConvertDataJSON(c.Response())
}

func (uc userController) UpdateUser(c echo.Context) error {
	user := models.CreateUserRequest{}
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	id, _ := strconv.Atoi(c.Param("id"))
	if err != nil {
		return wrapperResponse(http.StatusBadRequest, "Error", nil).ConvertDataJSON(c.Response())
	}

	newUser := &database.User{
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
	}
	errServices := uc.UserService.UpdateUser(newUser, id)
	if errServices != nil {
		return wrapperResponse(http.StatusInternalServerError, "Error", nil).ConvertDataJSON(c.Response())
	}
	return wrapperResponse(http.StatusOK, "Success update user", user.Email).ConvertDataJSON(c.Response())
}

func (uc userController) DeleteUser(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return wrapperResponse(http.StatusBadRequest, "Error", nil).ConvertDataJSON(c.Response())
	}

	errServices := uc.UserService.DeleteUser(id)
	if errServices != nil {
		return wrapperResponse(http.StatusInternalServerError, "Error", nil).ConvertDataJSON(c.Response())
	}
	return wrapperResponse(http.StatusOK, "Success Delete user", nil).ConvertDataJSON(c.Response())
}

func (uc userController) GetUserById(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return wrapperResponse(http.StatusBadRequest, "Error", nil).ConvertDataJSON(c.Response())
	}

	user, errServices := uc.UserService.GetUserById(id)
	if errServices != nil {
		return wrapperResponse(http.StatusInternalServerError, "Error", nil).ConvertDataJSON(c.Response())
	}
	return wrapperResponse(http.StatusOK, "Success get user", user).ConvertDataJSON(c.Response())
}
