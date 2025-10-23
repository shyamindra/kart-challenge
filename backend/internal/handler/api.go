package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"kart-challenge/backend/internal/model"
	"kart-challenge/backend/internal/service"
	"kart-challenge/backend/internal/storage"
)

// Server implements the ServerInterface for our API.
type Server struct {
	OrderService *service.OrderService
	ProductRepo  storage.ProductRepository
}

// ListProducts implements the /product endpoint.
func (s *Server) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := s.ProductRepo.GetAllProducts()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(products); err != nil {
		// Log error, but response is already started
		fmt.Printf("Error encoding products: %v\n", err)
	}
}

// GetProduct implements the /product/{productId} endpoint.
func (s *Server) GetProduct(w http.ResponseWriter, r *http.Request, productId int64) {
	// Convert int64 productId to string for repository lookup
	idStr := strconv.FormatInt(productId, 10)

	product, err := s.ProductRepo.GetProductByID(idStr)
	if err != nil {
		// Assuming the only error is "not found" for simplicity
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(product); err != nil {
		fmt.Printf("Error encoding product: %v\n", err)
	}
}

// PlaceOrder implements the /order endpoint.
func (s *Server) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var orderReq model.OrderReq
	if err := json.NewDecoder(r.Body).Decode(&orderReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	order, err := s.OrderService.PlaceOrder(orderReq)
	if err != nil {
		// Check for specific error types (e.g., invalid coupon, product not found)
		errStr := err.Error()

		if orderReq.CouponCode != nil && strings.Contains(errStr, "invalid coupon code") {
			http.Error(w, "Invalid promo code", http.StatusBadRequest)
			return
		}

		if strings.Contains(errStr, "product not found") {
			http.Error(w, "One or more products not found", http.StatusBadRequest)
			return
		}

		// Generic internal error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(order); err != nil {
		fmt.Printf("Error encoding order: %v\n", err)
	}
}
