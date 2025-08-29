package search

type SearchDocument struct {
	ID string `json:"id"`
	Type string `json:"type"` // "item" or "vendor"

	NameEn        string `json:"nameEn"`
	NameAr        string `json:"nameAr"`
	DescriptionEn string `json:"descriptionEn,omitempty"`
	DescriptionAr string `json:"descriptionAr,omitempty"`
	VendorNameEn  string `json:"vendorNameEn,omitempty"`
	VendorNameAr  string `json:"vendorNameAr,omitempty"`

	// --- Common and Pricing Fields ---
	Picture         string  `json:"picture"`
	BasePrice       int     `json:"basePrice,omitempty"`
	DiscountPercent float64 `json:"discountPercent"`

	// --- IDs and Filters ---
	VendorID   uint `json:"vendorId,omitempty"`
	DistrictID uint `json:"districtId"`
	CategoryID uint `json:"categoryId,omitempty"`
	IsActive   bool `json:"isActive"`

	// --- Vendor-specific field ---
	Items []string `json:"items,omitempty"`
	CoverPicture string `json:"coverPicture"`
}