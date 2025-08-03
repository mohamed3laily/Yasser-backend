package response

import (
	"net/http"
	"strings"
	"yasser-backend/pkg/i18n"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type APIResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
}

type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

type AppErrorInterface interface {
	Error() string
	GetStatusCode() int
	GetCode() string
}

func Success(c *gin.Context, messageKey string, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Message: i18n.T(c, messageKey),
		Data:    data,
	})
}

func OK(c *gin.Context, messageKey string) {
	c.JSON(http.StatusOK, APIResponse{
		Message: i18n.T(c, messageKey),
	})
}

func Created(c *gin.Context, messageKey string, data interface{}) {
	c.JSON(http.StatusCreated, APIResponse{
		Message: i18n.T(c, messageKey),
		Data:    data,
	})
}

func Error(c *gin.Context, appErr interface{}) {
	if err, ok := appErr.(AppErrorInterface); ok {
		c.JSON(err.GetStatusCode(), APIResponse{
			Message: i18n.T(c, "common.request_failed"),
			Error: &ErrorInfo{
				Code:    err.GetCode(),
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusInternalServerError, APIResponse{
		Message: i18n.T(c, "common.request_failed"),
		Error: &ErrorInfo{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: i18n.T(c, "common.internal_server_error"),
		},
	})
}

func ValidationError(c *gin.Context, err error) {
	var details string
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		details = formatValidationErrors(c, validationErrors)
	} else {
		details = err.Error()
	}

	c.JSON(http.StatusBadRequest, APIResponse{
		Message: i18n.T(c, "validation.failed"),
		Error: &ErrorInfo{
			Code:    "VALIDATION_ERROR",
			Message: i18n.T(c, "validation.invalid_request_data"),
			Details: details,
		},
	})
}

func formatValidationErrors(c *gin.Context, errs validator.ValidationErrors) string {
	var messages []string
	for _, err := range errs {
		var message string
		switch err.Tag() {
		case "required":
			message = i18n.T(c, "validation.required", err.Field())
		case "min":
			message = i18n.T(c, "validation.min", err.Field(), err.Param())
		case "max":
			message = i18n.T(c, "validation.max", err.Field(), err.Param())
		case "len":
			message = i18n.T(c, "validation.len", err.Field(), err.Param())
		case "numeric":
			message = i18n.T(c, "validation.numeric", err.Field())
		case "email":
			message = i18n.T(c, "validation.email", err.Field())
		case "url":
			message = i18n.T(c, "validation.url", err.Field())
		default:
			message = i18n.T(c, "validation.invalid", err.Field())
		}
		messages = append(messages, message)
	}

	return strings.Join(messages, ", ")
}