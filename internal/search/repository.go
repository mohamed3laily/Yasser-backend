package search

import (
	"fmt"
	"math"

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
	var results []ItemDTO

	err := r.db.Table("items i").
		Select(`i.id, i.name_en, i.name_ar, i.description_en, i.description_ar, 
			   i.picture ,i.base_price, i.discount_percent, i.vendor_id, i.category_id, i.is_active,
			   v.name_en as vendor_name_en, v.name_ar as vendor_name_ar, v.district_id`).
		Joins("LEFT JOIN vendors v ON v.id = i.vendor_id").
		Where("i.is_active = ?", true).
		Scan(&results).Error

	if err != nil {
		return nil, fmt.Errorf("failed to fetch items for indexing: %w", err)
	}

	documents := make([]SearchDocument, 0, len(results))
	fmt.Println("Items length: ", len(results))
	
	for _, item := range results {
		doc := SearchDocument{
			ID:              fmt.Sprintf("item_%d", item.ID),
			Type:            "item",
			NameEn:          item.NameEn,
			NameAr:          item.NameAr,
			DescriptionEn:   item.DescriptionEn,
			DescriptionAr:   item.DescriptionAr,
			VendorNameEn:    item.VendorNameEn,
			VendorNameAr:    item.VendorNameAr,
			Picture:         item.Picture,
			BasePrice:       item.BasePrice,
			DiscountPercent: item.DiscountPercent,
			VendorID:        item.VendorID,
			CategoryID:      item.CategoryID,
			DistrictID:      item.DistrictID,
			IsActive:        item.IsActive,
		}
		documents = append(documents, doc)
	}

	return documents, nil
}

func (r *repository) GetAllVendorsForIndexing() ([]SearchDocument, error) {
	var results []VendorDTO

	err := r.db.Table("vendors").
		Select("id, name_en, name_ar, description_en, description_ar, profile_picture, cover_picture, category_id, is_active, district_id").
		Where("is_active = ?", true).
		Scan(&results).Error

	if err != nil {
		return nil, fmt.Errorf("failed to fetch vendors for indexing: %w", err)
	}

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
			ID:            fmt.Sprintf("vendor_%d", vendor.ID),
			Type:          "vendor",
			NameEn:        vendor.NameEn,
			NameAr:        vendor.NameAr,
			DescriptionEn: vendor.DescriptionEn,
			DescriptionAr: vendor.DescriptionAr,
			Picture:       vendor.ProfilePicture,
			CoverPicture:  vendor.CoverPicture,
			CategoryID:    vendor.CategoryID,
			Items:         itemNames,
			DistrictID:    vendor.DistrictID,
			IsActive:      vendor.IsActive,
		}
		documents = append(documents, doc)
	}

	return documents, nil
}

func (r *repository) calculateFinalPrice(basePrice int, discountPercent float64) int {
	// If no discount or invalid discount, return base price
	if discountPercent <= 0 || basePrice <= 0 {
		return basePrice
	}
	
	// Calculate discounted price
	discountMultiplier := 1.0 - (discountPercent / 100.0)
	discountedPrice := float64(basePrice) * discountMultiplier
	
	// Ensure we don't return a price below 0
	finalPrice := int(math.Round(discountedPrice))
	if finalPrice < 0 {
		return 0
	}
	
	return finalPrice
}
