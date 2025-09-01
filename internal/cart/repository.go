package cart

import (
	"fmt"
	"yasser-backend/internal/item-group/item"

	"gorm.io/gorm"
)

type Repository interface {
	ValidateCartItems(vendorID int64, cartItems []CartItem) ([]CartItemValidation, error)
	GetItemWithRelations(itemID int64) (*item.Item, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
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

func (r *repository) ValidateCartItems(vendorID int64, cartItems []CartItem) ([]CartItemValidation, error) {
	var validations []CartItemValidation
	
	for _, cartItem := range cartItems {
		validation := r.validateSingleCartItem(vendorID, cartItem)
		validations = append(validations, validation)
	}
	
	return validations, nil
}

func (r *repository) validateSingleCartItem(vendorID int64, cartItem CartItem) CartItemValidation {
	validation := CartItemValidation{
		ItemID:  cartItem.ItemID,
		IsValid: true,
		Errors:  []string{},
	}
	
	// Get item with relations
	itemEntity, err := r.GetItemWithRelations(cartItem.ItemID)
	if err != nil {
		return r.createInvalidValidation(cartItem.ItemID, "cart.item_not_found")
	}
	
	// Validate vendor
	if itemEntity.VendorID != vendorID {
		return r.createInvalidValidation(cartItem.ItemID, "cart.item_vendor_mismatch")
	}
	
	// Perform all validations
	r.validateStock(itemEntity, cartItem, &validation)
	r.validateSize(itemEntity, cartItem, &validation)
	r.validateVariant(itemEntity, cartItem, &validation)
	r.validateAddons(itemEntity, cartItem, &validation)
	
	// Set item details if valid
	if validation.IsValid {
		validation.ItemDetails = r.createValidatedItemInfo(itemEntity)
	}
	
	return validation
}

func (r *repository) createInvalidValidation(itemID int64, errorKey string) CartItemValidation {
	return CartItemValidation{
		ItemID:  itemID,
		IsValid: false,
		Errors:  []string{errorKey},
	}
}

func (r *repository) validateStock(itemEntity *item.Item, cartItem CartItem, validation *CartItemValidation) {
	if itemEntity.Stock < cartItem.Quantity {
		validation.IsValid = false
		validation.Errors = append(validation.Errors, "cart.insufficient_stock")
	}
}

func (r *repository) validateSize(itemEntity *item.Item, cartItem CartItem, validation *CartItemValidation) {
	if cartItem.SizeID == nil {
		return
	}
	
	if !r.isValidID(*cartItem.SizeID, itemEntity.Sizes) {
		validation.IsValid = false
		validation.Errors = append(validation.Errors, "cart.invalid_size")
	}
}

func (r *repository) validateVariant(itemEntity *item.Item, cartItem CartItem, validation *CartItemValidation) {
	if cartItem.VariantID == nil {
		return
	}
	
	if !r.isValidVariantID(*cartItem.VariantID, itemEntity.Variants) {
		validation.IsValid = false
		validation.Errors = append(validation.Errors, "cart.invalid_variant")
	}
}

func (r *repository) validateAddons(itemEntity *item.Item, cartItem CartItem, validation *CartItemValidation) {
	for _, addonID := range cartItem.AddonIDs {
		if !r.isValidAddonID(addonID, itemEntity.Addons) {
			validation.IsValid = false
			validation.Errors = append(validation.Errors, fmt.Sprintf("Invalid addon ID: %d", addonID))
		}
	}
}

func (r *repository) isValidID(targetID int64, sizes []item.ItemSize) bool {
	for _, size := range sizes {
		if int64(size.ID) == targetID {
			return true
		}
	}
	return false
}

func (r *repository) isValidVariantID(targetID int64, variants []item.ItemVariant) bool {
	for _, variant := range variants {
		if int64(variant.ID) == targetID {
			return true
		}
	}
	return false
}

func (r *repository) isValidAddonID(targetID int64, addons []item.ItemAddon) bool {
	for _, addon := range addons {
		if int64(addon.ID) == targetID {
			return true
		}
	}
	return false
}

func (r *repository) createValidatedItemInfo(itemEntity *item.Item) *ValidatedItemInfo {
	finalPrice := float64(itemEntity.BasePrice) * (1 - itemEntity.DiscountPercent/100.0)
	return &ValidatedItemInfo{
		ID:              int64(itemEntity.ID),
		Name:            itemEntity.NameEn, // You can add language logic here
		BasePrice:       itemEntity.BasePrice,
		DiscountPercent: itemEntity.DiscountPercent,
		FinalPrice:      finalPrice,
		IsActive:        itemEntity.IsActive,
		Stock:           itemEntity.Stock,
		VendorID:        itemEntity.VendorID,
	}
}