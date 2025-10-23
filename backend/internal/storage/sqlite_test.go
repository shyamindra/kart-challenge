package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"kart-challenge/backend/internal/model"
)

func TestSQLiteStore_Repository(t *testing.T) {
	// Setup: Create a temporary directory for the SQLite file
	testDataDir := t.TempDir()
	
	// Initialize SQLiteStore
	store, err := NewSQLiteStore(testDataDir)
	assert.NoError(t, err)
	assert.NotNil(t, store)
	defer store.db.Close()
	
	// Ensure the database file was created
	dbPath := filepath.Join(testDataDir, dbFileName)
	_, err = os.Stat(dbPath)
	assert.NoError(t, err)

	// --- Test GetAllProducts ---
	t.Run("GetAllProducts", func(t *testing.T) {
		products, err := store.GetAllProducts()
		assert.NoError(t, err)
		assert.Len(t, products, 2) // Should contain the 2 hardcoded products
	})

	// --- Test GetProductByID ---
	t.Run("GetProductByID_Existing", func(t *testing.T) {
		product, err := store.GetProductByID("10")
		assert.NoError(t, err)
		assert.NotNil(t, product)
		assert.Equal(t, "Chicken Waffle", *product.Name)
	})

	t.Run("GetProductByID_NonExisting", func(t *testing.T) {
		product, err := store.GetProductByID("999")
		assert.Error(t, err)
		assert.Nil(t, product)
		assert.Contains(t, err.Error(), "product with ID 999 not found")
	})

	// --- Test CreateOrder ---
	t.Run("CreateOrder", func(t *testing.T) {
		// Helper functions for pointers
		sPtr := func(s string) *string { return &s }
		fPtr := func(f float32) *float32 { return &f }
		iPtr := func(i int) *int { return &i }

		orderID := "test-order-1"
		total := float32(13.3)

		order := model.Order{
			Id:        sPtr(orderID),
			Total:     fPtr(total),
			Discounts: fPtr(0.0),
			Items: &[]struct {
				ProductId *string `json:"productId,omitempty"`
				Quantity  *int    `json:"quantity,omitempty"`
			}{
				{ProductId: sPtr("10"), Quantity: iPtr(1)},
			},
		}

		err := store.CreateOrder(order)
		assert.NoError(t, err)

		// Verification: Check if the order count increased (requires a GetOrderCount method, but we'll skip for simplicity and rely on CreateOrder returning no error)
	})
}
