package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"kart-challenge/backend/internal/handler"
)

func TestListProducts(t *testing.T) {
	// Create a new Chi router
	r := chi.NewRouter()

	// Our API operations
	var api handler.ServerInterface = &handler.Server{}

	// Register the API handlers with the router
	handler.HandlerFromMux(api, r)

	// Create a test HTTP server
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Make a request to the /product endpoint
	res, err := http.Get(ts.URL + "/product")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// Read the response body
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	assert.NoError(t, err)
	assert.Equal(t, "[]", string(body))
}
