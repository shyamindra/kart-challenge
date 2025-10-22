package storage

import (
	"kart-challenge/backend/internal/handler"
)

// ProductRepository defines the interface for product data operations.
type ProductRepository interface {
	GetAllProducts() ([]handler.Product, error)
	GetProductByID(id string) (*handler.Product, error)
}

// OrderRepository defines the interface for order data operations.
type OrderRepository interface {
	CreateOrder(order handler.Order) error
}
