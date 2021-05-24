package errs

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func NewNotFoundError(m string) *AppError {
	return &AppError{
		Message: m,
		Code:    http.StatusNotFound,
	}
}

func NewUnexpectedError(m string) *AppError {
	return &AppError{
		Message: m,
		Code:    http.StatusInternalServerError,
	}
}
