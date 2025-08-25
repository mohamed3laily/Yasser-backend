package search

type SearchDocument struct {
	ID          string   `json:"id"`
	Type        string   `json:"type"` // "item" or "vendor"
	NameEn      string   `json:"nameEn"`
	NameAr      string   `json:"nameAr"`
	Description string   `json:"description,omitempty"`
	Picture     string   `json:"picture,omitempty"`
	BasePrice   int      `json:"basePrice,omitempty"`
	VendorID    uint     `json:"vendorId,omitempty"`
	CityID      uint     `json:"cityId,omitempty"`
	VendorName  string   `json:"vendorName,omitempty"`
	CategoryID  uint     `json:"categoryId,omitempty"`
	Items       []string `json:"items,omitempty"` // For vendors: list of item names
	IsActive    bool     `json:"isActive"`
}

type SearchResponse struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Picture     string `json:"picture,omitempty"`
	BasePrice   int    `json:"basePrice,omitempty"`
	VendorID    uint   `json:"vendorId,omitempty"`
	VendorName  string `json:"vendorName,omitempty"`
	CategoryID  uint   `json:"categoryId,omitempty"`
	
}

type SearchRequest struct {
	Query string `json:"query" validate:"required,min=1"`
	Lang  string `json:"lang" validate:"omitempty,oneof=en ar"`
	Limit int    `json:"limit" validate:"omitempty,min=1,max=100"`
	Type  string `json:"type" validate:"omitempty,oneof=item vendor"`
}