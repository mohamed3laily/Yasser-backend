package cart

type CartValidationRequest struct {
	VendorID int64      `json:"vendorId" validate:"required"`
	Items    []CartItem `json:"items" validate:"required,min=1"`
}

type CartItemValidation struct {
	ItemID      int64               `json:"itemId"`
	IsValid     bool                `json:"isValid"`
	Errors      []string            `json:"errors,omitempty"`
	ItemDetails *ValidatedItemInfo  `json:"itemDetails,omitempty"`
}