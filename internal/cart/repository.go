package cart

import (
	"context"
	"yasser-backend/internal/item-group/item"

	"gorm.io/gorm"
)

type Repository interface {
	GetItemWithRelations(ctx context.Context, itemID int64) (*item.Item, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetItemWithRelations(ctx context.Context, itemID int64) (*item.Item, error) {
	var entity item.Item
	err := r.db.WithContext(ctx).
		Preload("Variants").
		Preload("Sizes").
		Preload("Addons").
		Where("id = ? AND is_active = ?", itemID, true).
		First(&entity).Error

	if err != nil {
		return nil, err
	}
	return &entity, nil
}
