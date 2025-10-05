package order

// import (
// 	"yasser-backend/internal/cart"
// 	"yasser-backend/pkg/errors"
// )

// type Service interface {
// 	CreateOrder(request CreateOrderRequest) (*CreateOrderResponse, error)
// }

// type service struct {
// 	repo     Repository
// 	cartSvc  cart.Service
// }

// func NewService(repo Repository, cartSvc cart.Service) Service {
// 	// return &service{
// 	// 	repo:    repo,
// 	// 	cartSvc: cartSvc,
// 	// }
// }

// // func (s *service) CreateOrder(request CreateOrderRequest) (*CreateOrderResponse, error) {
// // 	// First validate the cart
// // 	cartValidation := cart.CartValidationRequest{
// // 		VendorID: request.VendorID,
// // 		Items:    request.Items,
// // 	}

// // 	validationResult, err := s.cartSvc.ValidateCart(cartValidation)
// // 	if err != nil {
// // 		if err == errors.ErrInvalid {
// // 			return nil, errors.ErrInvalid
// // 		}
// // 		return nil, fmt.Errorf("cart validation failed: %w", err)
// // 	}

// // 	if !validationResult.IsValid {
// // 		return nil, errors.ErrInvalid
// // 	}

// // 	// Create order
// // 	order := &Order{
// // 		VendorID:     request.VendorID,
// // 		UserID:   	  request.UserID,
// // 		Status:       OrderStatusPending,
// // 		TotalAmount:  validationResult.TotalPrice + request.DeliveryFee,
// // 		DeliveryFee:  request.DeliveryFee,
// // 		Notes:        request.Notes,
// // 		DeliveryTime: request.DeliveryTime,
// // 		Items:        []OrderItem{},
// // 	}

// // 	// Convert cart items to order items
// // 	for _, cartItem := range request.Items {
// // 		orderItem, err := s.convertCartItemToOrderItem(cartItem)
// // 		if err != nil {
// // 			return nil, fmt.Errorf("failed to convert cart item %d: %w", cartItem.ItemID, err)
// // 		}
// // 		order.Items = append(order.Items, *orderItem)
// // 	}

// // 	// Save order
// // 	if err := s.repo.CreateOrder(order); err != nil {
// // 		return nil, err
// // 	}

// // 	return &CreateOrderResponse{
// // 		OrderID:     int64(order.ID),
// // 		TotalAmount: order.TotalAmount,
// // 		Status:      string(order.Status),
// // 		CreatedAt:   order.CreatedAt.Format(time.RFC3339),
// // 	}, nil
// // }

// func (s *service) convertCartItemToOrderItem(cartItem cart.CartItem) (*OrderItem, error) {
// 	// Get item details
// 	itemEntity, err := s.repo.GetItemWithRelations(cartItem.ItemID)
// 	if err != nil {
// 		return nil, errors.ErrNotFound
// 	}

// 	unitPrice := float64(itemEntity.BasePrice) * (1 - itemEntity.DiscountPercent/100.0)

// 	orderItem := &OrderItem{
// 		ItemID:     cartItem.ItemID,
// 		ItemName:   itemEntity.NameEn, // Add language logic as needed
// 		Quantity:   cartItem.Quantity,
// 		UnitPrice:  unitPrice,
// 		TotalPrice: unitPrice * float64(cartItem.Quantity),
// 		Addons:     []OrderItemAddon{},
// 	}

// 	// Handle size
// 	if cartItem.SizeID != nil {
// 		for _, size := range itemEntity.Sizes {
// 			if int64(size.ID) == *cartItem.SizeID {
// 				orderItem.SizeID = cartItem.SizeID
// 				orderItem.SizeName = &size.Name
// 				orderItem.SizePrice = &size.Price
// 				break
// 			}
// 		}
// 	}

// 	// Handle variant
// 	if cartItem.VariantID != nil {
// 		for _, variant := range itemEntity.Variants {
// 			if int64(variant.ID) == *cartItem.VariantID {
// 				orderItem.VariantID = cartItem.VariantID
// 				orderItem.VariantName = &variant.NameEn // Add language logic
// 				break
// 			}
// 		}
// 	}

// 	// Handle addons
// 	for _, addonID := range cartItem.AddonIDs {
// 		for _, addon := range itemEntity.Addons {
// 			if int64(addon.ID) == addonID {
// 				orderAddon := OrderItemAddon{
// 					AddonID:    addonID,
// 					AddonName:  addon.NameEn, // Add language logic
// 					AddonPrice: addon.Price,
// 					IsRemoval:  addon.IsRemoval,
// 				}
// 				orderItem.Addons = append(orderItem.Addons, orderAddon)
// 				break
// 			}
// 		}
// 	}

// 	return orderItem, nil
// }