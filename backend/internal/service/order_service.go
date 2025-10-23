package service

import (
	"fmt"

	"github.com/google/uuid"

	"kart-challenge/backend/internal/model"
	"kart-challenge/backend/internal/promocode"
	"kart-challenge/backend/internal/storage"
)

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

	// 1. Validate coupon code
	var discount float32
	if req.CouponCode != nil && *req.CouponCode != "" {
		if !s.promo.IsValid(*req.CouponCode) {
			return nil, fmt.Errorf("invalid coupon code: %s", *req.CouponCode)
		}
		// For now, a valid coupon just gives a flat 10% discount
		discount = 0.10 // 10% discount
	}

	var total float32
	orderItems := make([]struct {
		ProductId *string `json:"productId,omitempty"`
		Quantity  *int    `json:"quantity,omitempty"`
	}, 0)

	// 2. Fetch product details and calculate subtotal
	for _, itemReq := range req.Items {
		product, err := s.productRepo.GetProductByID(itemReq.ProductId)
		if err != nil {
			return nil, fmt.Errorf("product not found: %s", itemReq.ProductId)
		}
		total += *product.Price * float32(itemReq.Quantity)

		orderItems = append(orderItems, struct {
			ProductId *string `json:"productId,omitempty"`
			Quantity  *int    `json:"quantity,omitempty"`
		}{
			ProductId: sPtr(itemReq.ProductId),
			Quantity:  iPtr(itemReq.Quantity),
		})
	}

	// Apply discount
	discountsAmount := total * discount
	total -= discountsAmount

	// 5. Construct the final types.Order object
	newOrder := model.Order{
		Id:        sPtr(uuid.New().String()),
		Items:     &orderItems,
		Total:     fPtr(total),
		Discounts: fPtr(discountsAmount),
	}

	// 6. Save the order
	if err := s.orderRepo.CreateOrder(newOrder); err != nil {
		return nil, fmt.Errorf("failed to save order: %w", err)
	}

	return &newOrder, nil
}
