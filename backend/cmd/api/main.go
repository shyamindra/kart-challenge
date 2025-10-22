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
)

func main() {
	log.Print("Starting server...")

	// Initialize promo code validator
	dataDir := filepath.Join(".", "data") // Assuming data directory is at the root of backend-challenge
	validator := promocode.NewValidator(dataDir)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Our API operations
	var api handler.ServerInterface = &handler.Server{PromoCodeValidator: validator}

	handler.HandlerFromMux(api, r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
