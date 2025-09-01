package order

import (
	"yasser-backend/internal/item-group/item"

	"gorm.io/gorm"
)

type Repository interface {
	CreateOrder(order *Order) error
	GetItemWithRelations(itemID int64) (*item.Item, error)
	UpdateItemStock(itemID int64, quantity int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateOrder(order *Order) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Create order
		if err := tx.Create(order).Error; err != nil {
			return err
		}
		
		// Update stock for each item
		for _, orderItem := range order.Items {
			if err := r.updateItemStockTx(tx, orderItem.ItemID, orderItem.Quantity); err != nil {
				return err
			}
		}
		
		return nil
	})
}

func (r *repository) GetItemWithRelations(itemID int64) (*item.Item, error) {
	var itemEntity item.Item
	err := r.db.Preload("Variants").Preload("Sizes").Preload("Addons").
		Where("id = ? AND is_active = ?", itemID, true).
		First(&itemEntity).Error
	
	if err != nil {
		return nil, err
	}
	return &itemEntity, nil
}

func (r *repository) UpdateItemStock(itemID int64, quantity int) error {
	return r.db.Model(&item.Item{}).
		Where("id = ?", itemID).
		Update("stock", gorm.Expr("stock - ?", quantity)).Error
}

func (r *repository) updateItemStockTx(tx *gorm.DB, itemID int64, quantity int) error {
	return tx.Model(&item.Item{}).
		Where("id = ? AND stock >= ?", itemID, quantity).
		Update("stock", gorm.Expr("stock - ?", quantity)).Error
}