package cart

import (
	"yasser-backend/pkg/errors"
)

type Service interface {
	ValidateCart(request CartValidationRequest) (*CartValidationResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) ValidateCart(request CartValidationRequest) (*CartValidationResponse, error) {
	if len(request.Items) == 0 {
		return nil, errors.ErrInvalid
	}
	
	validations, err := s.repo.ValidateCartItems(request.VendorID, request.Items)
	if err != nil {
		return nil, err
	}
	
	response := &CartValidationResponse{
		IsValid:    true,
		TotalPrice: 0,
		Items:      validations,
		Errors:     []string{},
	}
	
	// Calculate total price and check overall validity
	for _, validation := range validations {
		if !validation.IsValid {
			response.IsValid = false
			continue
		}
		
		if validation.ItemDetails != nil {
			// Find the cart item to get quantity
			for _, cartItem := range request.Items {
				if cartItem.ItemID == validation.ItemID {
					itemTotal := validation.ItemDetails.FinalPrice * float64(cartItem.Quantity)
					
					// Add size price if applicable
					if cartItem.SizeID != nil {
						// You'd need to fetch size price and add it
					}
					
					// Add addon prices if applicable
					if len(cartItem.AddonIDs) > 0 {
						// You'd need to fetch addon prices and add them
					}
					
					response.TotalPrice += itemTotal
					break
				}
			}
		}
	}
	
	if !response.IsValid {
		return response, errors.ErrInvalid
	}
	
	return response, nil
}