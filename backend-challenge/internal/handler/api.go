package handler

import (
	"net/http"
)

// Server implements the ServerInterface for our API.
type Server struct {
}

// ListProducts implements the /product endpoint.
func (s *Server) ListProducts(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("[]")) // Return an empty JSON array for now
}

// PlaceOrder implements the /order endpoint.
func (s *Server) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// GetProduct implements the /product/{productId} endpoint.
func (s *Server) GetProduct(w http.ResponseWriter, r *http.Request, productId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}
