package handler

import (
	"encoding/json"
	"net/http"

	"kart-challenge/backend/internal/promocode"
)

// Server implements the ServerInterface for our API.
type Server struct {
	PromoCodeValidator *promocode.Validator
}

// ListProducts implements the /product endpoint.
func (s *Server) ListProducts(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("[]")) // Return an empty JSON array for now
}

// PlaceOrder implements the /order endpoint.
func (s *Server) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var orderReq OrderReq
	if err := json.NewDecoder(r.Body).Decode(&orderReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if orderReq.CouponCode != nil && *orderReq.CouponCode != "" {
		if !s.PromoCodeValidator.IsValid(*orderReq.CouponCode) {
			http.Error(w, "Invalid promo code", http.StatusBadRequest)
			return
		}
	}

	// For now, just return a success status
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Order placed successfully!"))
}

// GetProduct implements the /product/{productId} endpoint.
func (s *Server) GetProduct(w http.ResponseWriter, r *http.Request, productId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}
