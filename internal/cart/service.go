package cart

import (
	"context"
	"fmt"
	"math"
	"yasser-backend/internal/item-group/item"
	"yasser-backend/pkg/errors"
)

type Service interface {
	ValidateCart(ctx context.Context, req CartValidationRequest) (*CartValidationResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) ValidateCart(ctx context.Context, req CartValidationRequest) (*CartValidationResponse, error) {
	if len(req.Items) == 0 {
		return nil, errors.ErrInvalid
	}

	response := s.initializeResponse()
	vendorID := s.validateCartItems(ctx, req.Items, response)
	
	if vendorID == nil {
		response.IsValid = false
	}

	return response, nil
}

func (s *service) initializeResponse() *CartValidationResponse {
	return &CartValidationResponse{
		IsValid:         true,
		Items:           []CartItemValidation{},
		TotalPrice:      0,
		OutOfStockIDs:   []int64{},
		PriceChangedIDs: []int64{},
	}
}

func (s *service) validateCartItems(ctx context.Context, items []CartItem, response *CartValidationResponse) *int64 {
	var vendorID *int64

	for i, cartItem := range items {
		itemEntity, err := s.repo.GetItemWithRelations(ctx, cartItem.ItemID)
		if err != nil {
			s.markItemAsInvalid(response, cartItem.ItemID, "cart.item_not_found")
			continue
		}

		if !s.checkVendorConsistency(i, itemEntity.VendorID, &vendorID, response, cartItem.ItemID) {
			continue
		}

		s.validateAndProcessItem(cartItem, itemEntity, response)
	}

	return vendorID
}

func (s *service) checkVendorConsistency(index int, currentVendorID int64, vendorID **int64, response *CartValidationResponse, itemID int64) bool {
	if index == 0 {
		*vendorID = &currentVendorID
		return true
	}

	if *vendorID != nil && currentVendorID != **vendorID {
		s.markItemAsInvalid(response, itemID, "cart.items_different_vendors")
		return false
	}

	return true
}

func (s *service) validateAndProcessItem(cartItem CartItem, itemEntity *item.Item, response *CartValidationResponse) {
	validation := s.buildItemValidation(cartItem, itemEntity)
	s.checkStock(cartItem, itemEntity, &validation, response)
	s.checkPrice(cartItem, itemEntity, &validation, response)

	if validation.IsValid {
		validation.ItemDetails = s.createItemDetails(itemEntity)
		s.addToTotal(cartItem, itemEntity, response)
	} else {
		response.IsValid = false
	}

	response.Items = append(response.Items, validation)
}

func (s *service) buildItemValidation(cartItem CartItem, itemEntity *item.Item) CartItemValidation {
	validation := CartItemValidation{
		ItemID:  cartItem.ItemID,
		IsValid: true,
		Errors:  []string{},
	}

	s.validateSize(cartItem, itemEntity, &validation)
	s.validateVariant(cartItem, itemEntity, &validation)
	s.validateAddons(cartItem, itemEntity, &validation)

	return validation
}

func (s *service) validateSize(cartItem CartItem, itemEntity *item.Item, validation *CartItemValidation) {
	if cartItem.SizeID != nil && !s.sizeExists(*cartItem.SizeID, itemEntity.Sizes) {
		validation.IsValid = false
		validation.Errors = append(validation.Errors, "cart.invalid_size")
	}
}

func (s *service) validateVariant(cartItem CartItem, itemEntity *item.Item, validation *CartItemValidation) {
	if cartItem.VariantID != nil && !s.variantExists(*cartItem.VariantID, itemEntity.Variants) {
		validation.IsValid = false
		validation.Errors = append(validation.Errors, "cart.invalid_variant")
	}
}

func (s *service) validateAddons(cartItem CartItem, itemEntity *item.Item, validation *CartItemValidation) {
	for _, addonID := range cartItem.AddonIDs {
		if !s.addonExists(addonID, itemEntity.Addons) {
			validation.IsValid = false
			validation.Errors = append(validation.Errors, fmt.Sprintf("cart.invalid_addon_%d", addonID))
		}
	}
}

func (s *service) checkStock(cartItem CartItem, itemEntity *item.Item, validation *CartItemValidation, response *CartValidationResponse) {
	if itemEntity.Stock < cartItem.Quantity {
		response.OutOfStockIDs = append(response.OutOfStockIDs, cartItem.ItemID)
		validation.IsValid = false
		validation.Errors = append(validation.Errors, "cart.insufficient_stock")
	}
}

func (s *service) checkPrice(cartItem CartItem, itemEntity *item.Item, validation *CartItemValidation, response *CartValidationResponse) {
	serverPrice := s.calculateItemPrice(cartItem, itemEntity)
	
	if !s.pricesMatch(serverPrice, cartItem.FinalPrice) {
		response.PriceChangedIDs = append(response.PriceChangedIDs, cartItem.ItemID)
		validation.IsValid = false
		validation.Errors = append(validation.Errors, "cart.price_changed")
	}
}

func (s *service) pricesMatch(price1, price2 float64) bool {
	return math.Abs(price1-price2) < 0.01
}

func (s *service) addToTotal(cartItem CartItem, itemEntity *item.Item, response *CartValidationResponse) {
	unitPrice := s.calculateItemPrice(cartItem, itemEntity)
	itemTotal := unitPrice * float64(cartItem.Quantity)
	response.TotalPrice += itemTotal
}

func (s *service) calculateItemPrice(cartItem CartItem, itemEntity *item.Item) float64 {
	price := s.calculateBasePrice(itemEntity)
	price += s.getSizePrice(cartItem.SizeID, itemEntity.Sizes)
	price += s.getAddonsPrice(cartItem.AddonIDs, itemEntity.Addons)
	return price
}

func (s *service) calculateBasePrice(itemEntity *item.Item) float64 {
	return float64(itemEntity.BasePrice) * (1 - itemEntity.DiscountPercent/100)
}

func (s *service) getSizePrice(sizeID *int64, sizes []item.ItemSize) float64 {
	if sizeID == nil {
		return 0
	}

	for _, sz := range sizes {
		if int64(sz.ID) == *sizeID {
			return float64(sz.Price)
		}
	}
	return 0
}

func (s *service) getAddonsPrice(addonIDs []int64, addons []item.ItemAddon) float64 {
	totalPrice := 0.0
	
	for _, addonID := range addonIDs {
		for _, addon := range addons {
			if int64(addon.ID) == addonID {
				totalPrice += float64(addon.Price)
				break
			}
		}
	}
	
	return totalPrice
}

func (s *service) markItemAsInvalid(response *CartValidationResponse, itemID int64, errorKey string) {
	response.IsValid = false
	response.Items = append(response.Items, CartItemValidation{
		ItemID:  itemID,
		IsValid: false,
		Errors:  []string{errorKey},
	})
}

func (s *service) createItemDetails(itemEntity *item.Item) *ValidatedItemInfo {
	finalPrice := s.calculateBasePrice(itemEntity)
	return &ValidatedItemInfo{
		ID:              int64(itemEntity.ID),
		Name:            itemEntity.NameEn,
		BasePrice:       itemEntity.BasePrice,
		DiscountPercent: itemEntity.DiscountPercent,
		FinalPrice:      finalPrice,
		IsActive:        itemEntity.IsActive,
		Stock:           itemEntity.Stock,
		VendorID:        itemEntity.VendorID,
	}
}

func (s *service) sizeExists(id int64, sizes []item.ItemSize) bool {
	for _, sz := range sizes {
		if int64(sz.ID) == id {
			return true
		}
	}
	return false
}

func (s *service) variantExists(id int64, variants []item.ItemVariant) bool {
	for _, v := range variants {
		if int64(v.ID) == id {
			return true
		}
	}
	return false
}

func (s *service) addonExists(id int64, addons []item.ItemAddon) bool {
	for _, a := range addons {
		if int64(a.ID) == id {
			return true
		}
	}
	return false
}