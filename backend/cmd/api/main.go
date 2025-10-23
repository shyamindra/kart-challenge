package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"kart-challenge/backend/internal/handler"
	"kart-challenge/backend/internal/promocode"
	"kart-challenge/backend/internal/service"
	"kart-challenge/backend/internal/storage"
)

func main() {
	log.Print("Starting server...")

	// Initialize promo code validator
	dataDir := filepath.Join("backend", "data")
	validator := promocode.NewValidator(dataDir)

	// Initialize storage
	store := storage.NewInMemoryStorage()

	// Initialize service
	orderService := service.NewOrderService(store, store, validator)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Our API operations
	var api handler.ServerInterface = &handler.Server{
		OrderService: orderService,
		ProductRepo:  store,
	}

	handler.HandlerFromMux(api, r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
