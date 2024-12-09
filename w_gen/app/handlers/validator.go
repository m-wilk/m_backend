package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type dataValidator struct {
	validator *validator.Validate
}

func NewValidator() *dataValidator {
	return &dataValidator{validator: validator.New()}
}

func (d *dataValidator) Validate(i interface{}) error {
	if err := d.validator.Struct(i); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ApiError, len(ve))
			for i, fe := range ve {
				out[i] = ApiError{fe.Field(), msgForTag(fe)}
			}
			return echo.NewHTTPError(http.StatusBadRequest, out)
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

type ApiError struct {
	Param   string `json:"param"`
	Message string `json:"massage"`
}

func msgForTag(fe validator.FieldError) string {
	// TODO - more case statements
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "min":
		return fmt.Sprintf("minimum value %s", fe.Param())
	}

	return "unexpected validation error"
}
