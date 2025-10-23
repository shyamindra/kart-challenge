package storage

import (
	"kart-challenge/backend/internal/model"
)

// ProductRepository defines the interface for product data operations.
type ProductRepository interface {
	GetAllProducts() ([]model.Product, error)
	GetProductByID(id string) (*model.Product, error)
}

// OrderRepository defines the interface for order data operations.
type OrderRepository interface {
	CreateOrder(order model.Order) error
}

// Store combines all repository interfaces into a single interface for dependency injection.
type Store interface {
	ProductRepository
	OrderRepository
}
