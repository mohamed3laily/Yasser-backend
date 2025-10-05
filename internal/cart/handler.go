package cart

import (
	"yasser-backend/pkg/errors"
	"yasser-backend/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service   Service
	validator *validator.Validate
}

func NewHandler(service Service, validator *validator.Validate) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}

func (h *Handler) ValidateCart(c *gin.Context) {
	var request CartValidationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.ValidationError(c, err)
		return
	}

	if err := h.validator.Struct(&request); err != nil {
		response.ValidationError(c, err)
		return
	}

	validationResult, err := h.service.ValidateCart(c.Request.Context(), request)
	if err != nil {
		appErr := errors.Handle(c, err, "cart.validation_failed")
		response.Error(c, appErr)
		return
	}

	if validationResult.IsValid {
		response.Success(c, "cart.validation_success", validationResult)
	} else {
		response.Success(c, "cart.validation_failed", validationResult)
	}
}