package errors

import (
	"errors"
	"net/http"
	"yasser-backend/pkg/i18n"

	"github.com/gin-gonic/gin"
)

var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrInvalid      = errors.New("invalid")
	ErrExpired      = errors.New("expired")
)

type AppError struct {
	MessageKey string `json:"-"`   
	StatusCode int    `json:"-"`    
	Code       string `json:"code"`    
	Params     []interface{} `json:"-"`
	context    *gin.Context  `json:"-"`
}

func (e *AppError) Error() string {
	if e.context != nil {
		return i18n.T(e.context, e.MessageKey, e.Params...)
	}
	return i18n.TWithLang("en", e.MessageKey, e.Params...)
}

func (e *AppError) GetStatusCode() int {
	return e.StatusCode
}

func (e *AppError) GetCode() string {
	return e.Code
}

func (e *AppError) WithContext(c *gin.Context) *AppError {
	e.context = c
	return e
}

func BadRequest(messageKey string, params ...interface{}) *AppError {
	return &AppError{
		MessageKey: messageKey,
		StatusCode: http.StatusBadRequest,
		Code:       "BAD_REQUEST",
		Params:     params,
	}
}

func NotFound(messageKey string, params ...interface{}) *AppError {
	return &AppError{
		MessageKey: messageKey,
		StatusCode: http.StatusNotFound,
		Code:       "NOT_FOUND",
		Params:     params,
	}
}

func Unauthorized(messageKey string, params ...interface{}) *AppError {
	return &AppError{
		MessageKey: messageKey,
		StatusCode: http.StatusUnauthorized,
		Code:       "UNAUTHORIZED",
		Params:     params,
	}
}

func Internal(messageKey string, params ...interface{}) *AppError {
	return &AppError{
		MessageKey: messageKey,
		StatusCode: http.StatusInternalServerError,
		Code:       "INTERNAL_SERVER_ERROR",
		Params:     params,
	}
}

func Handle(c *gin.Context, err error, fallbackMessageKey string) *AppError {
	if err == nil {
		return nil
	}

	var appErr *AppError
	switch {
	case errors.Is(err, ErrNotFound):
		appErr = NotFound("common.not_found")
	case errors.Is(err, ErrUnauthorized):
		appErr = Unauthorized("common.unauthorized")
	case errors.Is(err, ErrInvalid):
		appErr = BadRequest("common.bad_request")
	case errors.Is(err, ErrExpired):
		appErr = Unauthorized("auth.token_expired")
	default:
		appErr = Internal(fallbackMessageKey)
	}

	return appErr.WithContext(c)
}