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
	VendorName  string `json:"vendorName,omitempty"`
	Picture     string `json:"picture,omitempty"`
	CoverPicture   string `json:"coverPicture"`

	// Pricing
	BasePrice       int     `json:"basePrice,omitempty"`
	DiscountPercent float64 `json:"discountPercent,omitempty"`
	Price           int     `json:"price,omitempty"`

	// IDs
	VendorID   uint `json:"vendorId,omitempty"`
	CategoryID uint `json:"categoryId,omitempty"`
}

type ItemDTO struct {
	ID              uint    `json:"id"`
	NameEn          string  `json:"nameEn"`
	NameAr          string  `json:"nameAr"`
	DescriptionEn   string  `json:"descriptionEn"`
	DescriptionAr   string  `json:"descriptionAr"`
	Picture         string  `json:"picture"`
	BasePrice       int     `json:"basePrice"`
	DiscountPercent float64 `json:"discountPercent"`
	VendorID        uint    `json:"vendorId"`
	CategoryID      uint    `json:"categoryId"`
	IsActive        bool    `json:"isActive"`
	VendorNameEn    string  `json:"vendorNameEn"`
	VendorNameAr    string  `json:"vendorNameAr"`
	DistrictID      uint    `json:"districtId"`
}

type VendorDTO struct {
	ID             uint   `json:"id"`
	NameEn         string `json:"nameEn"`
	NameAr         string `json:"nameAr"`
	DescriptionEn  string `json:"descriptionEn"`
	DescriptionAr  string `json:"descriptionAr"`
	ProfilePicture string `json:"profilePicture"`
	CoverPicture   string `json:"coverPicture"`
	CategoryID     uint   `json:"categoryId"`
	IsActive       bool   `json:"isActive"`
	DistrictID     uint   `json:"districtId"`
}