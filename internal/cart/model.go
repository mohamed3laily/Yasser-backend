package cart

import "github.com/lib/pq"


type CartItem struct {
	ItemID    int64   `json:"itemId" validate:"required"`
	Quantity  int     `json:"quantity" validate:"required,min=1"`
	SizeID    *int64  `json:"sizeId,omitempty"`
	VariantID *int64  `json:"variantId,omitempty"`
	AddonIDs  pq.Int64Array  `gorm:"type:bigint[]" json:"addonIds,omitempty"`
}

type ValidatedItemInfo struct {
	ID              int64   `json:"id"`
	Name            string  `json:"name"`
	BasePrice       int     `json:"basePrice"`
	DiscountPercent float64 `json:"discountPercent"`
	FinalPrice      float64 `json:"finalPrice"`
	IsActive        bool    `json:"isActive"`
	Stock           int     `json:"stock"`
	VendorID        int64   `json:"vendorId"`
}