package middleware

import (
	"errors"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func NewCustomValidator() echo.Validator {
	return &CustomValidator{validator: validator.New()}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return errors.New(strconv.Itoa(http.StatusBadRequest))
	}
	return nil
}
