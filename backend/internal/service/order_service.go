package service

import (
	"fmt"

	"github.com/google/uuid"

	"kart-challenge/backend/internal/model"
	"kart-challenge/backend/internal/promocode"
	"kart-challenge/backend/internal/storage"
)

// OrderServiceIface defines the interface for the OrderService.
type OrderServiceIface interface {
	PlaceOrder(req model.OrderReq) (*model.Order, error)
}

// OrderService defines the business logic for orders.
type OrderService struct {
	productRepo storage.ProductRepository
	orderRepo   storage.OrderRepository
	promo       promocode.PromoCodeValidator
}

// NewOrderService creates a new OrderService.
func NewOrderService(productRepo storage.ProductRepository, orderRepo storage.OrderRepository, promo promocode.PromoCodeValidator) *OrderService {
	return &OrderService{
		productRepo: productRepo,
		orderRepo:   orderRepo,
		promo:       promo,
	}
}

// PlaceOrder processes a new order request.
func (s *OrderService) PlaceOrder(req model.OrderReq) (*model.Order, error) {
	// Helper functions for pointers
	sPtr := func(s string) *string { return &s }
	fPtr := func(f float32) *float32 { return &f }
	iPtr := func(i int) *int { return &i }

	// Internal struct to hold item details and price for calculation
	type pricedItem struct {
		productID string
		quantity  int
		price     float32
		subtotal  float32
	}

	var total float32
	var discountsAmount float32
	pricedItems := make([]pricedItem, 0, len(req.Items))

	// 1. Fetch product details and calculate subtotal
	for _, itemReq := range req.Items {
		product, err := s.productRepo.GetProductByID(itemReq.ProductId)
		if err != nil {
			return nil, fmt.Errorf("product not found: %s", itemReq.ProductId)
		}

		itemPrice := *product.Price
		itemSubtotal := itemPrice * float32(itemReq.Quantity)
		total += itemSubtotal

		pricedItems = append(pricedItems, pricedItem{
			productID: itemReq.ProductId,
			quantity:  itemReq.Quantity,
			price:     itemPrice,
			subtotal:  itemSubtotal,
		})
	}

	// 2. Validate and apply coupon code
	if req.CouponCode != nil && *req.CouponCode != "" {
		coupon := *req.CouponCode
		if !s.promo.IsValid(coupon) {
			return nil, fmt.Errorf("invalid coupon code: %s", coupon)
		}

		switch coupon {
		case "HAPPYHOURS":
			// 20% discount on total
			discountsAmount = total * 0.20
		case "BUYGETONE":
			// Lowest priced item is free
			if len(pricedItems) > 0 {
				lowestPrice := pricedItems[0].price
				for _, item := range pricedItems {
					if item.price < lowestPrice {
						lowestPrice = item.price
					}
				}
				discountsAmount = lowestPrice
			}
		default:
			// Default 10% discount for any other valid coupon
			discountsAmount = total * 0.10
		}
	}

	// Apply discount
	total -= discountsAmount

	// 3. Construct the final types.Order object
	orderItems := make([]struct {
		ProductId *string `json:"productId,omitempty"`
		Quantity  *int    `json:"quantity,omitempty"`
	}, len(pricedItems))

	for i, item := range pricedItems {
		orderItems[i] = struct {
			ProductId *string `json:"productId,omitempty"`
			Quantity  *int    `json:"quantity,omitempty"`
		}{
			ProductId: sPtr(item.productID),
			Quantity:  iPtr(item.quantity),
		}
	}

	newOrder := model.Order{
		Id:        sPtr(uuid.New().String()),
		Items:     &orderItems,
		Total:     fPtr(total),
		Discounts: fPtr(discountsAmount),
	}

	// 4. Save the order
	if err := s.orderRepo.CreateOrder(newOrder); err != nil {
		return nil, fmt.Errorf("failed to save order: %w", err)
	}

	return &newOrder, nil
}
