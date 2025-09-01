package cart

import (
	"strconv"
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
	vendorIDStr := c.Param("vendorId")
	vendorID, err := strconv.ParseInt(vendorIDStr, 10, 64)
	if err != nil {
		appErr := errors.BadRequest("cart.invalid_vendor_id").WithContext(c)
		response.Error(c, appErr)
		return
	}
	
	var request CartValidationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.ValidationError(c, err)
		return
	}
	
	request.VendorID = vendorID
	
	if err := h.validator.Struct(&request); err != nil {
		response.ValidationError(c, err)
		return
	}
	
	validationResult, err := h.service.ValidateCart(request)
	if err != nil {
		appErr := errors.Handle(c, err, "cart.validation_failed")
		response.Error(c, appErr)
		return
	}
	
	response.Success(c, "cart.validation_success", validationResult)
}