package storage

import (
	"fmt"
	"sync"

	"kart-challenge/backend/internal/handler"
)

// InMemoryStorage implements ProductRepository and OrderRepository using in-memory data.
type InMemoryStorage struct {
	products map[string]handler.Product
	orders   []handler.Order
	mu       sync.RWMutex
}

// NewInMemoryStorage creates and initializes a new InMemoryStorage.
func NewInMemoryStorage() *InMemoryStorage {
	s := &InMemoryStorage{
		products: make(map[string]handler.Product),
		orders:   []handler.Order{},
	}
	s.initProducts()
	return s
}

// Helper function to get a pointer to a string
func sPtr(s string) *string { return &s }

// Helper function to get a pointer to a float32
func fPtr(f float32) *float32 { return &f }

// createProduct is a helper to create handler.Product instances with correct pointer fields.
func createProduct(id, name, category string, price float32, desktop, mobile, thumbnail, tablet string) handler.Product {
	return handler.Product{
		Id:       sPtr(id),
		Name:     sPtr(name),
		Price:    fPtr(price),
		Category: sPtr(category),
		Image: &struct {
			Desktop   *string `json:"desktop,omitempty"`
			Mobile    *string `json:"mobile,omitempty"`
			Tablet    *string `json:"tablet,omitempty"`
			Thumbnail *string `json:"thumbnail,omitempty"`
		}{
			Desktop:   sPtr(desktop),
			Mobile:    sPtr(mobile),
			Thumbnail: sPtr(thumbnail),
			Tablet:    sPtr(tablet),
		},
	}
}

// initProducts hardcodes the initial product data.
func (s *InMemoryStorage) initProducts() {
	s.products["10"] = createProduct(
		"10",
		"Chicken Waffle",
		"Waffle",
		13.3,
		"https://orderfoodonline.deno.dev/public/images/image-waffle-desktop.jpg",
		"https://orderfoodonline.deno.dev/public/images/image-waffle-mobile.jpg",
		"https://orderfoodonline.deno.dev/public/images/image-waffle-thumbnail.jpg",
		"https://orderfoodonline.deno.dev/public/images/image-waffle-tablet.jpg",
	)
	s.products["11"] = createProduct(
		"11",
		"Beef Burger",
		"Burger",
		15.0,
		"https://orderfoodonline.deno.dev/public/images/image-burger-desktop.jpg",
		"https://orderfoodonline.deno.dev/public/images/image-burger-mobile.jpg",
		"https://orderfoodonline.deno.dev/public/images/image-burger-thumbnail.jpg",
		"https://orderfoodonline.deno.dev/public/images/image-burger-tablet.jpg",
	)
	// Add more products as needed
}

// GetAllProducts returns all available products.
func (s *InMemoryStorage) GetAllProducts() ([]handler.Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	products := make([]handler.Product, 0, len(s.products))
	for _, p := range s.products {
		products = append(products, p)
	}
	return products, nil
}

// GetProductByID returns a product by its ID.
func (s *InMemoryStorage) GetProductByID(id string) (*handler.Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	p, ok := s.products[id]
	if !ok {
		return nil, fmt.Errorf("product with ID %s not found", id)
	}
	return &p, nil
}

// CreateOrder adds a new order to the storage.
func (s *InMemoryStorage) CreateOrder(order handler.Order) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.orders = append(s.orders, order)
	return nil
}