package item

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	GetByID(id uint) (*Item, error)
	GetByVendorID(vendorID uint) ([] ItemWithCategory , error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetByID(id uint) (*Item, error) {
	var item Item
	err := r.db.Preload("Variants").Preload("Sizes").Preload("Addons").
		First(&item, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &item, nil
}

func (r *repository) GetByVendorID(vendorID uint) ([]ItemWithCategory, error) {
	var items []ItemWithCategory
	err := r.db.
		Table("items").
		Select(`items.id, items.name_en, items.name_ar, items.base_price, items.discount, items.vendor_id,
		        items.picture, items.category_id, items.description_ar, items.description_en, c.name_en as category_name_en, c.name_ar as category_name_ar`).
		Joins("JOIN items_categories c ON c.id = items.category_id").
		Where("items.vendor_id = ? AND items.is_active = TRUE", vendorID).
		Scan(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}