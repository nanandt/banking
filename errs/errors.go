package errs

import "net/http"

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
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

func NewValidationError(m string) *AppError {
	return &AppError{
		Message: m,
		Code:    http.StatusUnprocessableEntity,
	}
}
