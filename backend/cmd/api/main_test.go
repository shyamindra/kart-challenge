package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"kart-challenge/backend/internal/handler"
	"kart-challenge/backend/internal/model"
	"kart-challenge/backend/internal/service"
)

// MockProductRepository is a mock implementation of storage.ProductRepository
type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) GetAllProducts() ([]model.Product, error) {
	args := m.Called()
	return args.Get(0).([]model.Product), args.Error(1)
}

func (m *MockProductRepository) GetProductByID(id string) (*model.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Product), args.Error(1)
}

func TestListProducts(t *testing.T) {
	// Setup mocks
	mockProductRepo := new(MockProductRepository)
	
	// Expected empty response
	mockProductRepo.On("GetAllProducts").Return([]model.Product{}, nil).Once()

	// Create a new Chi router
	r := chi.NewRouter()

	// Our API operations - inject mock dependencies
	var api handler.ServerInterface = &handler.Server{
		// OrderService is not used in ListProducts, but must be initialized to satisfy the struct
		OrderService: &service.OrderService{}, 
		ProductRepo:  mockProductRepo,
	}

	// Register the API handlers with the router
	handler.HandlerFromMux(api, r)

	// Create a test HTTP server
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Make a request to the /product endpoint
	res, err := http.Get(ts.URL + "/product")
	assert.NoError(t, err)
	defer res.Body.Close()
	
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// Read the response body
	body, err := io.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, "[]\n", string(body))
	
	mockProductRepo.AssertExpectations(t)
}