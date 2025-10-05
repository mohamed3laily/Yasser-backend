package cart

import "github.com/lib/pq"

type CartValidationRequest struct {
	VendorID   int64      `json:"vendorId" validate:"required"`
	Items      []CartItem `json:"items" validate:"required,min=1,dive"`
	TotalPrice float64    `json:"totalPrice" validate:"required,gt=0"`
}

type CartItem struct {
	ItemID    int64         `json:"itemId" validate:"required"`
	Quantity  int           `json:"quantity" validate:"required,min=1"`
	SizeID    *int64        `json:"sizeId,omitempty"`
	VariantID *int64        `json:"variantId,omitempty"`
	AddonIDs  pq.Int64Array `gorm:"type:bigint[]" json:"addonIds,omitempty"`

	FinalPrice float64 `json:"finalPrice" validate:"required,gt=0"`
	Subtotal   float64 `json:"subtotal" validate:"required,gte=0"`
}

type CartItemValidation struct {
	ItemID      int64              `json:"itemId"`
	IsValid     bool               `json:"isValid"`
	Errors      []string           `json:"errors,omitempty"`
	ItemDetails *ValidatedItemInfo `json:"itemDetails,omitempty"`
}