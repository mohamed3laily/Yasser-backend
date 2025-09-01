package cart
type CartValidationResponse struct {
	IsValid    bool                 `json:"isValid"`
	TotalPrice float64              `json:"totalPrice"`
	Items      []CartItemValidation `json:"items"`
	Errors     []string             `json:"errors,omitempty"`
}