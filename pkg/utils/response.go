package utils

import (
	"net/http"
	"time"

	"github.com/cpartogi/withdrawdeposit/schema/response"
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

	responseData := response.Base{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    message,
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Info("success response")

	return ctx.JSON(http.StatusOK, responseData)
}

// CreatedResponse returns
func CreatedResponse(ctx echo.Context, message string, data interface{}) error {

	responseData := response.Base{
		Status:     "success insert data",
		StatusCode: http.StatusCreated,
		Message:    message,
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Info("success create data")

	return ctx.JSON(http.StatusCreated, responseData)
}

// ErrorResponse returns
func ErrorResponse(ctx echo.Context, err error, data interface{}) error {
	statusCode, err := errorType(err)
	switch statusCode {
	case http.StatusConflict:
		return ErrorConflictResponse(ctx, err, data)
	case http.StatusBadRequest:
		return ErrorBadRequest(ctx, err, data)
	case http.StatusNotFound:
		return ErrorNotFound(ctx, err, data)
	}
	return ErrorInternalServerResponse(ctx, err, data)
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

// ErrorInternalServerResponse returns
func ErrorInternalServerResponse(ctx echo.Context, err error, data interface{}) error {

	responseData := response.Base{
		Status:     "internal server error",
		StatusCode: http.StatusInternalServerError,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Errorf("internal server error : %s ", err.Error())

	return ctx.JSON(http.StatusInternalServerError, responseData)
}

// ErrorBadRequest returns
func ErrorBadRequest(ctx echo.Context, err error, data interface{}) error {
	responseData := response.Base{
		Status:     "bad request",
		StatusCode: http.StatusBadRequest,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Errorf("bad request error : %s ", err.Error())

	return ctx.JSON(http.StatusBadRequest, responseData)
}

// ErrorNotFound returns
func ErrorNotFound(ctx echo.Context, err error, data interface{}) error {
	responseData := response.Base{
		Status:     "not found",
		StatusCode: http.StatusNotFound,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Errorf("error not found : %s ", err.Error())

	return ctx.JSON(http.StatusNotFound, responseData)
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
	message := switchErrorValidation(err)
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
