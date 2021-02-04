package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cpartogi/withdrawdeposit/schema/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	log "go.uber.org/zap"
)

//ParsingError is
type ParsingError struct {
	msg string
}

func (re *ParsingError) Error() string { return re.msg }

// SuccessResponse returns
func SuccessResponse(ctx echo.Context, message string, data interface{}) error {

	responseData := response.SuccessResponse{
		Base: response.Base{
			Status:     "success",
			StatusCode: http.StatusOK,
			Message:    message,
			Timestamp:  time.Now().UTC(),
		},
		Data: data,
	}

	log.S().Info("success response")

	return ctx.JSON(http.StatusOK, responseData)
}

// CreatedResponse returns
func CreatedResponse(ctx echo.Context, message string, data interface{}) error {

	responseData := response.SuccessResponse{
		Base: response.Base{
			Status:     "success insert data",
			StatusCode: http.StatusCreated,
			Message:    message,
			Timestamp:  time.Now().UTC(),
		},
		Data: data,
	}

	log.S().Info("success create data")

	return ctx.JSON(http.StatusCreated, responseData)
}

// ErrorConflictResponse returns
func ErrorConflictResponse(ctx echo.Context, err error, data interface{}) error {

	responseData := response.Base{
		Status:     "conflict",
		StatusCode: http.StatusConflict,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Errorf("conflict data error : %s ", err.Error())

	return ctx.JSON(http.StatusConflict, responseData)
}

// ErrorParsing returns
func ErrorParsing(ctx echo.Context, err error, data interface{}) error {

	responseData := response.Base{
		Status:     "unprocessable entity",
		StatusCode: http.StatusUnprocessableEntity,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Errorf("parsing data error : %s ", err.Error())

	return ctx.JSON(http.StatusUnprocessableEntity, responseData)
}

// ErrorValidate returns
func ErrorValidate(ctx echo.Context, err error, data interface{}) error {
	var message string

	message = switchErrorValidation(err)
	responseData := response.Base{
		Status:     "bad request",
		StatusCode: http.StatusBadRequest,
		Message:    message,
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Errorf("validate data error : %s ", err.Error())

	return ctx.JSON(http.StatusBadRequest, responseData)
}

// ErrorParsingValidate returns
func ErrorParsingValidate(ctx echo.Context, err error, data interface{}) (errs error) {
	switch err.(type) {
	default:
		errs = ErrorValidate(ctx, err, data)
	case *ParsingError:
		errs = ErrorParsing(ctx, err, data)
	}

	return errs
}

// FailedReturn returns
func FailedReturn(ctx echo.Context, msg string, err error, data interface{}) error {
	responseData := response.Base{
		Status:     "bad request",
		StatusCode: http.StatusBadRequest,
		Message:    msg,
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Errorf("failed data error : %s ", err)

	return ctx.JSON(http.StatusBadRequest, responseData)
}

// InternalError returns
func InternalError(ctx echo.Context, msg string, err error, data interface{}) error {
	responseData := response.Base{
		Status:     "internal server error",
		StatusCode: http.StatusInternalServerError,
		Message:    msg,
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Errorf("internal error : %s ", err.Error())

	return ctx.JSON(http.StatusInternalServerError, responseData)
}

// UnprocessableEntity returns
func UnprocessableEntity(ctx echo.Context, msg string, err error, data interface{}) error {
	responseData := response.Base{
		Status:     "unprocessable entity",
		StatusCode: http.StatusUnprocessableEntity,
		Message:    msg,
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Errorf("error : %s ", err.Error())

	return ctx.JSON(http.StatusUnprocessableEntity, responseData)
}

// BadRequest returns
func BadRequest(ctx echo.Context, msg string, err error) error {
	responseData := response.Base{
		Status:     "bad request",
		StatusCode: http.StatusBadRequest,
		Message:    msg,
		Timestamp:  time.Now().UTC(),
	}

	log.S().Errorf("error : %s ", err.Error())

	return ctx.JSON(http.StatusBadRequest, responseData)
}

func switchErrorValidation(err error) (message string) {
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			field := SetLowerAndAddSpace(err.Field())

			// Change Field Name
			switch field {
			case "msisdn":
				field = "phone number"
			}

			// Check Error Type
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("%s is mandatory",
					field)
			case "email":
				message = fmt.Sprintf("%s is not valid email",
					field)
			case "number":
				message = fmt.Sprintf("%s must be numbers only",
					field)
			case "gte":
				message = fmt.Sprintf("%s value must be greater than %s",
					field, err.Param())
			case "lte":
				message = fmt.Sprintf("%s value must be lower than %s",
					field, err.Param())
			case "min":
				minimum := calculateMin(field, err.Param())
				message = fmt.Sprintf("%s at least %s characters long",
					field, minimum)
			case "startswith":
				message = fmt.Sprintf("%s must starts with %s",
					field, err.Param())
			}
			break
		}
	}
	return
}
