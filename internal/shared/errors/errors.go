package errors

import (
	baseErrors "errors"

	"github.com/labstack/echo/v4"
)

var (
	ErrNotFound        error = baseErrors.New("content not found")
	ErrRequestNotValid error = baseErrors.New("request not valid, please check the fields")
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

func NewHttpAppError(code int, message string, err error) *echo.HTTPError {
	appError := AppError{
		Code:    code,
		Message: message,
		Details: err.Error(),
	}

	return echo.NewHTTPError(code, appError)
}
