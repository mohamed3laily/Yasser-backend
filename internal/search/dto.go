package search


type SearchRequest struct {
	Query      string `validate:"required"`
	Lang       string `validate:"omitempty,oneof=en ar"`
	Type       *string `validate:"omitempty,oneof=item vendor"`

	DistrictID uint `validate:"required,number"`

	Limit  int `validate:"omitempty,min=1,max=100"`
	Offset int `validate:"omitempty,min=0"`
}

type SearchResponse struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	VendorName  string `json:"vendorName,omitempty"` // <-- ADDED
	Picture     string `json:"picture,omitempty"`

	// Pricing
	BasePrice       int     `json:"basePrice,omitempty"`
	DiscountPercent float64 `json:"discountPercent,omitempty"`
	DiscountedPrice int     `json:"discountedPrice,omitempty"`

	// IDs
	VendorID   uint `json:"vendorId,omitempty"`
	CategoryID uint `json:"categoryId,omitempty"`
}