package cart

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