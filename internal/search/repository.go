package search

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	GetAllItemsForIndexing() ([]SearchDocument, error)
	GetAllVendorsForIndexing() ([]SearchDocument, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAllItemsForIndexing() ([]SearchDocument, error) {
	var results []struct {
		ID            uint    `json:"id"`
		NameEn        string  `json:"nameEn"`
		NameAr        string  `json:"nameAr"`
		DescriptionEn string  `json:"descriptionEn"`
		DescriptionAr string  `json:"descriptionAr"`
		Picture       string  `json:"picture"`
		BasePrice     int     `json:"basePrice"`
		VendorID      uint    `json:"vendorId"`
		CategoryID    uint    `json:"categoryId"`
		IsActive      bool    `json:"isActive"`
		VendorNameEn  string  `json:"vendorNameEn"`
		VendorNameAr  string  `json:"vendorNameAr"`
	}

	err := r.db.Table("items i").
		Select(`i.id, i.name_en, i.name_ar, i.description_en, i.description_ar, 
			   i.picture, i.base_price, i.vendor_id, i.category_id, i.is_active,
			   v.name_en as vendor_name_en, v.name_ar as vendor_name_ar`).
		Joins("LEFT JOIN vendors v ON v.id = i.vendor_id").
		Where("i.is_active = ?", true).
		Scan(&results).Error

	if err != nil {
		return nil, fmt.Errorf("failed to fetch items for indexing: %w", err)
	}

	documents := make([]SearchDocument, 0, len(results))
	for _, item := range results {
		doc := SearchDocument{
			ID:         fmt.Sprintf("item_%d", item.ID),
			Type:       "item",
			NameEn:     item.NameEn,
			NameAr:     item.NameAr,
			Description: item.DescriptionEn + " " + item.DescriptionAr,
			Picture:    item.Picture,
			BasePrice:  item.BasePrice,
			VendorID:   item.VendorID,
			VendorName: item.VendorNameEn + " " + item.VendorNameAr,
			CategoryID: item.CategoryID,
			IsActive:   item.IsActive,
		}
		documents = append(documents, doc)
	}

	return documents, nil
}

func (r *repository) GetAllVendorsForIndexing() ([]SearchDocument, error) {
	var results []struct {
		ID             uint   `json:"id"`
		NameEn         string `json:"nameEn"`
		NameAr         string `json:"nameAr"`
		DescriptionEn  string `json:"descriptionEn"`
		DescriptionAr  string `json:"descriptionAr"`
		ProfilePicture string `json:"profilePicture"`
		CategoryID     uint   `json:"categoryId"`
		IsActive       bool   `json:"isActive"`
	}

	// First get vendors
	err := r.db.Table("vendors").
		Select("id, name_en, name_ar, description_en, description_ar, profile_picture, category_id, is_active").
		Where("is_active = ?", true).
		Scan(&results).Error

	if err != nil {
		return nil, fmt.Errorf("failed to fetch vendors for indexing: %w", err)
	}

	// Then get items for each vendor
	documents := make([]SearchDocument, 0, len(results))
	for _, vendor := range results {
		var itemNames []string
		err := r.db.Table("items").
			Where("vendor_id = ? AND is_active = ?", vendor.ID, true).
			Pluck("CONCAT(name_en, ' ', name_ar)", &itemNames).Error

		if err != nil {
			return nil, fmt.Errorf("failed to fetch item names for vendor %d: %w", vendor.ID, err)
		}

		doc := SearchDocument{
			ID:          fmt.Sprintf("vendor_%d", vendor.ID),
			Type:        "vendor",
			NameEn:      vendor.NameEn,
			NameAr:      vendor.NameAr,
			Description: vendor.DescriptionEn + " " + vendor.DescriptionAr,
			Picture:     vendor.ProfilePicture,
			CategoryID:  vendor.CategoryID,
			Items:       itemNames,
			IsActive:    vendor.IsActive,
		}
		documents = append(documents, doc)
	}

	return documents, nil
}