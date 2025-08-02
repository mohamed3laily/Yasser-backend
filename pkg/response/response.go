package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
}

type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, appErr interface{}) {
	if err, ok := appErr.(AppErrorInterface); ok {
		c.JSON(err.GetStatusCode(), APIResponse{
			Success: false,
			Message: "Request failed",
			Error: &ErrorInfo{
				Code:    err.GetCode(),
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusInternalServerError, APIResponse{
		Success: false,
		Message: "Request failed",
		Error: &ErrorInfo{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "An unexpected error occurred",
		},
	})
}

func ValidationError(c *gin.Context, err error) {
	var details string
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		details = formatValidationErrors(validationErrors)
	} else {
		details = err.Error()
	}

	c.JSON(http.StatusBadRequest, APIResponse{
		Success: false,
		Message: "Validation failed",
		Error: &ErrorInfo{
			Code:    "VALIDATION_ERROR",
			Message: "Invalid request data",
			Details: details,
		},
	})
}

type AppErrorInterface interface {
	Error() string
	GetStatusCode() int
	GetCode() string
}

func formatValidationErrors(errs validator.ValidationErrors) string {
	var messages []string
	for _, err := range errs {
		switch err.Tag() {
		case "required":
			messages = append(messages, err.Field()+" is required")
		case "min":
			messages = append(messages, err.Field()+" must be at least "+err.Param()+" characters")
		case "max":
			messages = append(messages, err.Field()+" must be at most "+err.Param()+" characters")
		case "len":
			messages = append(messages, err.Field()+" must be exactly "+err.Param()+" characters")
		case "numeric":
			messages = append(messages, err.Field()+" must contain only numbers")
		default:
			messages = append(messages, err.Field()+" is invalid")
		}
	}
	
	if len(messages) == 1 {
		return messages[0]
	}
	
	result := messages[0]
	for i := 1; i < len(messages); i++ {
		result += ", " + messages[i]
	}
	return result
}