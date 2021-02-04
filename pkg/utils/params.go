package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// ParsingAndValidateParameter will parsing request to struct and validate
func ParsingAndValidateParameter(ctx echo.Context, i interface{}) error {
	err := ctx.Bind(i)
	if err != nil {
		return &ParsingError{err.Error()}
	}

	validate := validator.New()
	err = validate.Struct(i)

	return err
}

// ParsingParameter will parsing request to struct
func ParsingParameter(ctx echo.Context, i interface{}) error {
	err := ctx.Bind(i)
	if err != nil {
		return &ParsingError{err.Error()}
	}
	return err
}

// ValidateParameter will validate request
func ValidateParameter(ctx echo.Context, i interface{}) (err error) {
	validate := validator.New()
	err = validate.Struct(i)

	return
}
