package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"kart-challenge/backend/internal/handler"
)

// Helper functions for pointers
func sPtr(s string) *string { return &s }
func fPtr(f float32) *float32 { return &f }

// MockProductRepository is a mock implementation of storage.ProductRepository
type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) GetAllProducts() ([]handler.Product, error) {
	args := m.Called()
	return args.Get(0).([]handler.Product), args.Error(1)
}

func (m *MockProductRepository) GetProductByID(id string) (*handler.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*handler.Product), args.Error(1)
}

// MockOrderRepository is a mock implementation of storage.OrderRepository
type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) CreateOrder(order handler.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

// MockPromoValidator is a mock implementation of promocode.Validator
type MockPromoValidator struct {
	mock.Mock
}

func (m *MockPromoValidator) IsValid(code string) bool {
	args := m.Called(code)
	return args.Bool(0)
}

func TestOrderService_PlaceOrder(t *testing.T) {
	// Helper functions for pointers
	sPtr := func(s string) *string { return &s }
	fPtr := func(f float32) *float32 { return &f }

	// Define a sample product (can be outside t.Run as it's constant)
	sampleProduct := handler.Product{
		Id:       sPtr("10"),
		Name:     sPtr("Chicken Waffle"),
		Price:    fPtr(13.3),
		Category: sPtr("Waffle"),
				Image: &struct {
					Desktop   *string `json:"desktop,omitempty"`
					Mobile    *string `json:"mobile,omitempty"`
					Tablet    *string `json:"tablet,omitempty"`
					Thumbnail *string `json:"thumbnail,omitempty"`
				}{
					Desktop:   nil,
					Mobile:    nil,
					Tablet:    nil,
					Thumbnail: nil,
				},
			}

	// Test case 1: Successful order without coupon
	t.Run("Successful order without coupon", func(t *testing.T) {
		// Setup mocks for this subtest
		mockProductRepo := new(MockProductRepository)
		mockOrderRepo := new(MockOrderRepository)
		mockPromoValidator := new(MockPromoValidator)
		service := NewOrderService(mockProductRepo, mockOrderRepo, mockPromoValidator)

		orderReq := handler.OrderReq{
			Items: []struct {
				ProductId string `json:"productId"`
				Quantity  int    `json:"quantity"`
			}{
				{ProductId: "10", Quantity: 2},
			},
		}

		mockProductRepo.On("GetProductByID", "10").Return(&sampleProduct, nil).Once()
		mockOrderRepo.On("CreateOrder", mock.AnythingOfType("handler.Order")).Return(nil).Once()
		mockPromoValidator.AssertNotCalled(t, "IsValid", mock.Anything)

		order, err := service.PlaceOrder(orderReq)
		assert.NoError(t, err)
		assert.NotNil(t, order)
		assert.Equal(t, float32(26.6), *order.Total)
		assert.Equal(t, float32(0), *order.Discounts)

		mockProductRepo.AssertExpectations(t)
		mockOrderRepo.AssertExpectations(t)
		mockPromoValidator.AssertExpectations(t)
	})

	// Test case 2: Successful order with valid coupon
	t.Run("Successful order with valid coupon", func(t *testing.T) {
		// Setup mocks for this subtest
		mockProductRepo := new(MockProductRepository)
		mockOrderRepo := new(MockOrderRepository)
		mockPromoValidator := new(MockPromoValidator)
		service := NewOrderService(mockProductRepo, mockOrderRepo, mockPromoValidator)

		couponCode := "VALIDCODE"
		orderReq := handler.OrderReq{
			CouponCode: sPtr(couponCode),
			Items: []struct {
				ProductId string `json:"productId"`
				Quantity  int    `json:"quantity"`
			}{
				{ProductId: "10", Quantity: 2},
			},
		}

		mockProductRepo.On("GetProductByID", "10").Return(&sampleProduct, nil).Once()
		mockOrderRepo.On("CreateOrder", mock.AnythingOfType("handler.Order")).Return(nil).Once()
		mockPromoValidator.On("IsValid", couponCode).Return(true).Once()

		order, err := service.PlaceOrder(orderReq)
		assert.NoError(t, err)
		assert.NotNil(t, order)
		// Expected total: 26.6 - (26.6 * 0.10) = 26.6 - 2.66 = 23.94
		assert.InDelta(t, float32(23.94), *order.Total, 0.001)
		assert.InDelta(t, float32(2.66), *order.Discounts, 0.001)

		mockProductRepo.AssertExpectations(t)
		mockOrderRepo.AssertExpectations(t)
		mockPromoValidator.AssertExpectations(t)
	})

	// Test case 3: Order with invalid coupon
	t.Run("Order with invalid coupon", func(t *testing.T) {
		// Setup mocks for this subtest
		mockProductRepo := new(MockProductRepository)
		mockOrderRepo := new(MockOrderRepository)
		mockPromoValidator := new(MockPromoValidator)
		service := NewOrderService(mockProductRepo, mockOrderRepo, mockPromoValidator)

		couponCode := "INVALIDCODE"
		orderReq := handler.OrderReq{
			CouponCode: sPtr(couponCode),
			Items: []struct {
				ProductId string `json:"productId"`
				Quantity  int    `json:"quantity"`
			}{
				{ProductId: "10", Quantity: 1},
			},
		}

		mockProductRepo.On("GetProductByID", "10").Return(&sampleProduct, nil).Maybe()
		mockPromoValidator.On("IsValid", couponCode).Return(false).Once()

		order, err := service.PlaceOrder(orderReq)
		assert.Error(t, err)
		assert.Nil(t, order)
		assert.Contains(t, err.Error(), "invalid coupon code")

		mockProductRepo.AssertExpectations(t)
		mockOrderRepo.AssertExpectations(t)
		mockPromoValidator.AssertExpectations(t)
	})

	// Test case 4: Product not found
	t.Run("Product not found", func(t *testing.T) {
		// Setup mocks for this subtest
		mockProductRepo := new(MockProductRepository)
		mockOrderRepo := new(MockOrderRepository)
		mockPromoValidator := new(MockPromoValidator)
		service := NewOrderService(mockProductRepo, mockOrderRepo, mockPromoValidator)

		orderReq := handler.OrderReq{
			Items: []struct {
				ProductId string `json:"productId"`
				Quantity  int    `json:"quantity"`
			}{
				{ProductId: "NONEXISTENT", Quantity: 1},
			},
		}

		mockProductRepo.On("GetProductByID", "NONEXISTENT").Return(nil, errors.New("product not found")).Once()

		order, err := service.PlaceOrder(orderReq)
		assert.Error(t, err)
		assert.Nil(t, order)
		assert.Contains(t, err.Error(), "product not found")

		mockProductRepo.AssertExpectations(t)
		mockOrderRepo.AssertNotCalled(t, "CreateOrder")
		mockPromoValidator.AssertNotCalled(t, "IsValid")
	})

	// Test case 5: Error saving order
	t.Run("Error saving order", func(t *testing.T) {
		// Setup mocks for this subtest
		mockProductRepo := new(MockProductRepository)
		mockOrderRepo := new(MockOrderRepository)
		mockPromoValidator := new(MockPromoValidator)
		service := NewOrderService(mockProductRepo, mockOrderRepo, mockPromoValidator)

		orderReq := handler.OrderReq{
			Items: []struct {
				ProductId string `json:"productId"`
				Quantity  int    `json:"quantity"`
			}{
				{ProductId: "10", Quantity: 1},
			},
		}

		mockProductRepo.On("GetProductByID", "10").Return(&sampleProduct, nil).Once()
		mockOrderRepo.On("CreateOrder", mock.AnythingOfType("handler.Order")).Return(errors.New("db error")).Once()
		mockPromoValidator.AssertNotCalled(t, "IsValid", mock.Anything)

		order, err := service.PlaceOrder(orderReq)
		assert.Error(t, err)
		assert.Nil(t, order)
		assert.Contains(t, err.Error(), "failed to save order")

		mockProductRepo.AssertExpectations(t)
		mockOrderRepo.AssertExpectations(t)
		mockPromoValidator.AssertExpectations(t)
	})
}
