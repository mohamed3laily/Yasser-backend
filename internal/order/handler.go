package order

// import (
// 	"yasser-backend/pkg/errors"
// 	"yasser-backend/pkg/response"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/validator/v10"
// )

// type Handler struct {
// 	service   Service
// 	validator *validator.Validate
// }

// func NewHandler(service Service, validator *validator.Validate) *Handler {
// 	return &Handler{
// 		service:   service,
// 		validator: validator,
// 	}
// }

// func (h *Handler) CreateOrder(c *gin.Context) {
// 	var request CreateOrderRequest
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		response.ValidationError(c, err)
// 		return
// 	}

// 	if err := h.validator.Struct(&request); err != nil {
// 		response.ValidationError(c, err)
// 		return
// 	}

// 	orderResponse, err := h.service.CreateOrder(request)
// 	if err != nil {
// 		appErr := errors.Handle(c, err, "order.creation_failed")
// 		response.Error(c, appErr)
// 		return
// 	}

// 	response.Created(c, "order.created_success", orderResponse)
// }