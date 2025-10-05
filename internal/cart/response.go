package cart
type CartValidationResponse struct {
	IsValid    bool                 `json:"isValid"`
	TotalPrice float64              `json:"totalPrice"`
	Items      []CartItemValidation `json:"items"`

	OutOfStockIDs  []int64 `json:"outOfStockIds,omitempty"`
	PriceChangedIDs []int64 `json:"priceChangedIds,omitempty"`
}