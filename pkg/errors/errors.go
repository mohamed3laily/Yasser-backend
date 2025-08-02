package errors

import (
	"errors"
	"net/http"
)

var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrInvalid      = errors.New("invalid")
	ErrExpired      = errors.New("expired")
)

type AppError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
	Code       string `json:"code"`
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) GetStatusCode() int {
	return e.StatusCode
}

func (e *AppError) GetCode() string {
	return e.Code
}

// Generic HTTP error constructors
func BadRequest(message string) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Code:       "BAD_REQUEST",
	}
}

func NotFound(message string) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: http.StatusNotFound,
		Code:       "NOT_FOUND",
	}
}

func Unauthorized(message string) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: http.StatusUnauthorized,
		Code:       "UNAUTHORIZED",
	}
}

func Internal(message string) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
		Code:       "INTERNAL_SERVER_ERROR",
	}
}

func Handle(err error, fallbackMessage string) *AppError {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, ErrNotFound):
		return NotFound(fallbackMessage)
	case errors.Is(err, ErrUnauthorized):
		return Unauthorized(fallbackMessage)
	case errors.Is(err, ErrInvalid):
		return BadRequest(fallbackMessage)
	case errors.Is(err, ErrExpired):
		return Unauthorized(fallbackMessage)
	default:
		return Internal(fallbackMessage)
	}
}