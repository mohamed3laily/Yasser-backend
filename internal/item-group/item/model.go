package item

import (
	itemcategory "yasser-backend/internal/item-group/item-category"
	"yasser-backend/pkg/models"
)

type Item struct {
	models.BaseModel
	NameEn        string    `gorm:"type:text;not null" json:"nameEn"`
	NameAr        string    `gorm:"type:text;not null" json:"nameAr"`
	DescriptionEn string    `gorm:"type:text" json:"descriptionEn"`
	DescriptionAr string    `gorm:"type:text" json:"descriptionAr"`
	BasePrice     int       `gorm:"type:int;not null" json:"basePrice"`
	Picture       string    `gorm:"type:text" json:"picture"`
	CategoryID    int64     `gorm:"type:bigint;index;not null" json:"categoryId"`
	VendorID      int64     `gorm:"type:bigint;index;not null" json:"vendorId"`
	DiscountPercent  float64   `gorm:"type:decimal(5,2);default:0;check:discount_percent >= 0 AND discount_percent <= 100" json:"discountPercent"`
	Stock         int       `gorm:"type:int;default:0" json:"stock"`
	IsActive      bool      `gorm:"default:true" json:"isActive"`


	// Relations
	Category itemcategory.ItemsCategory `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Variants []ItemVariant `gorm:"foreignKey:ItemID" json:"variants,omitempty"`
	Sizes    []ItemSize    `gorm:"foreignKey:ItemID" json:"sizes,omitempty"`
	Addons   []ItemAddon   `gorm:"foreignKey:ItemID" json:"addons,omitempty"`
}

type ItemVariant struct {
	models.BaseModel
    NameAr  string  `gorm:"type:text" json:"nameAr"`
	NameEn  string  `gorm:"type:text" json:"nameEn"`
    ItemID int64 `gorm:"type:bigint;index" json:"itemId"`
	VendorID int64 `gorm:"type:bigint;index" json:"vendorId"`
}

type ItemSize struct {
	models.BaseModel
	Name      string `gorm:"type:text" json:"name"`
	Price    int       `gorm:"type:int;not null" json:"price"`
    ItemID int64 `gorm:"type:bigint;index" json:"itemId"`
	VendorID int64 `gorm:"type:bigint;index" json:"vendorId"`
}

type ItemAddon struct {
	models.BaseModel
	NameEn      string `gorm:"type:text" json:"nameEn"`
	NameAr    string `gorm:"type:text" json:"nameAr"`
    Price     int64 `gorm:"type:bigint" json:"price"`
    IsRemoval bool  `gorm:"type:boolean" json:"isRemoval"`
    ItemID    int64 `gorm:"type:bigint;index" json:"itemId"`
	VendorID int64 `gorm:"type:bigint;index" json:"vendorId"`
}