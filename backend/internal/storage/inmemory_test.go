package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"kart-challenge/backend/internal/handler"
)

func TestInMemoryStorage_GetAllProducts(t *testing.T) {
	s := NewInMemoryStorage()
	products, err := s.GetAllProducts()
	assert.NoError(t, err)
	assert.Len(t, products, 2) // Assuming 2 products are hardcoded
}

func TestInMemoryStorage_GetProductByID(t *testing.T) {
	s := NewInMemoryStorage()

	// Test existing product
	product, err := s.GetProductByID("10")
	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "Chicken Waffle", *product.Name)

	// Test non-existing product
	product, err = s.GetProductByID("999")
	assert.Error(t, err)
	assert.Nil(t, product)
}

func TestInMemoryStorage_CreateOrder(t *testing.T) {
	s := NewInMemoryStorage()
	initialOrderCount := len(s.orders)

	// Helper functions for pointers
	sPtr := func(s string) *string { return &s }
	fPtr := func(f float32) *float32 { return &f }
	iPtr := func(i int) *int { return &i }

	orderID := "test-order-1"
	productID := "10"
	quantity := 1
	total := float32(13.3)

	order := handler.Order{
		Id:    sPtr(orderID),
		Total: fPtr(total),
		Items: &[]struct {
			ProductId *string `json:"productId,omitempty"`
			Quantity  *int    `json:"quantity,omitempty"`
		}{
			{ProductId: sPtr(productID), Quantity: iPtr(quantity)},
		},
	}

	err := s.CreateOrder(order)
	assert.NoError(t, err)
	assert.Len(t, s.orders, initialOrderCount+1)
	assert.Equal(t, *order.Id, *s.orders[initialOrderCount].Id)
}
