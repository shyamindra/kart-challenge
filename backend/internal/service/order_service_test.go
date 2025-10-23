package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"kart-challenge/backend/internal/model"
)

// Helper functions for pointers
func sPtr(s string) *string   { return &s }
func fPtr(f float32) *float32 { return &f }

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

// MockOrderRepository is a mock implementation of storage.OrderRepository
type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) CreateOrder(order model.Order) error {
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

	// Define sample products
	productWaffle := model.Product{
		Id:       sPtr("10"),
		Name:     sPtr("Chicken Waffle"),
		Price:    fPtr(13.3),
		Category: sPtr("Waffle"),
		Image:    &struct{
			Desktop   *string `json:"desktop,omitempty"`
			Mobile    *string `json:"mobile,omitempty"`
			Tablet    *string `json:"tablet,omitempty"`
			Thumbnail *string `json:"thumbnail,omitempty"`
		}{},
	}
	productBurger := model.Product{
		Id:       sPtr("11"),
		Name:     sPtr("Beef Burger"),
		Price:    fPtr(15.0),
		Category: sPtr("Burger"),
		Image:    &struct{
			Desktop   *string `json:"desktop,omitempty"`
			Mobile    *string `json:"mobile,omitempty"`
			Tablet    *string `json:"tablet,omitempty"`
			Thumbnail *string `json:"thumbnail,omitempty"`
		}{},
	}

	// Test case 1: Successful order without coupon
	t.Run("Successful order without coupon", func(t *testing.T) {
		mockProductRepo := new(MockProductRepository)
		mockOrderRepo := new(MockOrderRepository)
		mockPromoValidator := new(MockPromoValidator)
		service := NewOrderService(mockProductRepo, mockOrderRepo, mockPromoValidator)

		orderReq := model.OrderReq{
			Items: []struct {
				ProductId string `json:"productId"`
				Quantity  int    `json:"quantity"`
			}{
				{ProductId: "10", Quantity: 2},
			},
		}

		mockProductRepo.On("GetProductByID", "10").Return(&productWaffle, nil).Once()
		mockOrderRepo.On("CreateOrder", mock.AnythingOfType("model.Order")).Return(nil).Once()

		order, err := service.PlaceOrder(orderReq)
		assert.NoError(t, err)
		assert.NotNil(t, order)
		assert.InDelta(t, float32(26.6), *order.Total, 0.001)
		assert.InDelta(t, float32(0), *order.Discounts, 0.001)

		mockProductRepo.AssertExpectations(t)
		mockOrderRepo.AssertExpectations(t)
		mockPromoValidator.AssertNotCalled(t, "IsValid", mock.Anything)
	})

	// Test case 2: Successful order with default 10% coupon
	t.Run("Successful order with default 10% coupon", func(t *testing.T) {
		mockProductRepo := new(MockProductRepository)
		mockOrderRepo := new(MockOrderRepository)
		mockPromoValidator := new(MockPromoValidator)
		service := NewOrderService(mockProductRepo, mockOrderRepo, mockPromoValidator)

		couponCode := "DEFAULT10"
		orderReq := model.OrderReq{
			CouponCode: sPtr(couponCode),
			Items: []struct {
				ProductId string `json:"productId"`
				Quantity  int    `json:"quantity"`
			}{
				{ProductId: "10", Quantity: 2},
			},
		}

		mockProductRepo.On("GetProductByID", "10").Return(&productWaffle, nil).Once()
		mockOrderRepo.On("CreateOrder", mock.AnythingOfType("model.Order")).Return(nil).Once()
		mockPromoValidator.On("IsValid", couponCode).Return(true).Once()

		// Total: 26.6. Discount: 26.6 * 0.10 = 2.66. Final: 23.94
		order, err := service.PlaceOrder(orderReq)
		assert.NoError(t, err)
		assert.NotNil(t, order)
		assert.InDelta(t, float32(23.94), *order.Total, 0.001)
		assert.InDelta(t, float32(2.66), *order.Discounts, 0.001)

		mockProductRepo.AssertExpectations(t)
		mockOrderRepo.AssertExpectations(t)
		mockPromoValidator.AssertExpectations(t)
	})

	// Test case 3: Successful order with HAPPYHOURS (20%) coupon
	t.Run("Successful order with HAPPYHOURS (20%) coupon", func(t *testing.T) {
		mockProductRepo := new(MockProductRepository)
		mockOrderRepo := new(MockOrderRepository)
		mockPromoValidator := new(MockPromoValidator)
		service := NewOrderService(mockProductRepo, mockOrderRepo, mockPromoValidator)

		couponCode := "HAPPYHOURS"
		orderReq := model.OrderReq{
			CouponCode: sPtr(couponCode),
			Items: []struct {
				ProductId string `json:"productId"`
				Quantity  int    `json:"quantity"`
			}{
				{ProductId: "10", Quantity: 2},
			},
		}

		mockProductRepo.On("GetProductByID", "10").Return(&productWaffle, nil).Once()
		mockOrderRepo.On("CreateOrder", mock.AnythingOfType("model.Order")).Return(nil).Once()
		mockPromoValidator.On("IsValid", couponCode).Return(true).Once()

		// Total: 26.6. Discount: 26.6 * 0.20 = 5.32. Final: 21.28
		order, err := service.PlaceOrder(orderReq)
		assert.NoError(t, err)
		assert.NotNil(t, order)
		assert.InDelta(t, float32(21.28), *order.Total, 0.001)
		assert.InDelta(t, float32(5.32), *order.Discounts, 0.001)

		mockProductRepo.AssertExpectations(t)
		mockOrderRepo.AssertExpectations(t)
		mockPromoValidator.AssertExpectations(t)
	})

	// Test case 4: Successful order with BUYGETONE (lowest item free) coupon
	t.Run("Successful order with BUYGETONE coupon", func(t *testing.T) {
		mockProductRepo := new(MockProductRepository)
		mockOrderRepo := new(MockOrderRepository)
		mockPromoValidator := new(MockPromoValidator)
		service := NewOrderService(mockProductRepo, mockOrderRepo, mockPromoValidator)

		couponCode := "BUYGETONE"
		orderReq := model.OrderReq{
			CouponCode: sPtr(couponCode),
			Items: []struct {
				ProductId string `json:"productId"`
				Quantity  int    `json:"quantity"`
			}{
				{ProductId: "10", Quantity: 1}, // Price 13.3
				{ProductId: "11", Quantity: 1}, // Price 15.0
			},
		}

		mockProductRepo.On("GetProductByID", "10").Return(&productWaffle, nil).Once()
		mockProductRepo.On("GetProductByID", "11").Return(&productBurger, nil).Once()
		mockOrderRepo.On("CreateOrder", mock.AnythingOfType("model.Order")).Return(nil).Once()
		mockPromoValidator.On("IsValid", couponCode).Return(true).Once()

		// Subtotal: 13.3 + 15.0 = 28.3. Lowest price: 13.3. Discount: 13.3. Final: 15.0
		order, err := service.PlaceOrder(orderReq)
		assert.NoError(t, err)
		assert.NotNil(t, order)
		assert.InDelta(t, float32(15.0), *order.Total, 0.001)
		assert.InDelta(t, float32(13.3), *order.Discounts, 0.001)

		mockProductRepo.AssertExpectations(t)
		mockOrderRepo.AssertExpectations(t)
		mockPromoValidator.AssertExpectations(t)
	})

	// Test case 5: Order with invalid coupon
	t.Run("Order with invalid coupon", func(t *testing.T) {
		mockProductRepo := new(MockProductRepository)
		mockOrderRepo := new(MockOrderRepository)
		mockPromoValidator := new(MockPromoValidator)
		service := NewOrderService(mockProductRepo, mockOrderRepo, mockPromoValidator)

		couponCode := "INVALIDCODE"
		orderReq := model.OrderReq{
			CouponCode: sPtr(couponCode),
			Items: []struct {
				ProductId string `json:"productId"`
				Quantity  int    `json:"quantity"`
			}{
				{ProductId: "10", Quantity: 1},
			},
		}

		mockProductRepo.On("GetProductByID", "10").Return(&productWaffle, nil).Maybe()
		mockPromoValidator.On("IsValid", couponCode).Return(false).Once()

		order, err := service.PlaceOrder(orderReq)
		assert.Error(t, err)
		assert.Nil(t, order)
		assert.Contains(t, err.Error(), "invalid coupon code")

		mockProductRepo.AssertExpectations(t)
		mockOrderRepo.AssertExpectations(t)
		mockPromoValidator.AssertExpectations(t)
	})

	// Test case 6: Product not found
	t.Run("Product not found", func(t *testing.T) {
		mockProductRepo := new(MockProductRepository)
		mockOrderRepo := new(MockOrderRepository)
		mockPromoValidator := new(MockPromoValidator)
		service := NewOrderService(mockProductRepo, mockOrderRepo, mockPromoValidator)

		orderReq := model.OrderReq{
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

		mockProductRepo.AssertNotCalled(t, "CreateOrder")
		mockPromoValidator.AssertNotCalled(t, "IsValid")
	})

	// Test case 7: Error saving order
	t.Run("Error saving order", func(t *testing.T) {
		mockProductRepo := new(MockProductRepository)
		mockOrderRepo := new(MockOrderRepository)
		mockPromoValidator := new(MockPromoValidator)
		service := NewOrderService(mockProductRepo, mockOrderRepo, mockPromoValidator)

		orderReq := model.OrderReq{
			Items: []struct {
				ProductId string `json:"productId"`
				Quantity  int    `json:"quantity"`
			}{
				{ProductId: "10", Quantity: 1},
			},
		}

		mockProductRepo.On("GetProductByID", "10").Return(&productWaffle, nil).Once()
		mockOrderRepo.On("CreateOrder", mock.AnythingOfType("model.Order")).Return(errors.New("db error")).Once()
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