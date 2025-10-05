package order

import (
	"time"
	"yasser-backend/internal/cart"
)

type CreateOrderRequest struct {
	VendorID     int64            `json:"vendorId" validate:"required"`
	UserID   int64            	  `json:"userId" validate:"required"`
	Items        []cart.CartItem  `json:"items" validate:"required,min=1"`
	DeliveryFee  float64          `json:"deliveryFee,omitempty"`
	Notes        string           `json:"notes,omitempty"`
	DeliveryTime *time.Time       `json:"deliveryTime,omitempty"`
}