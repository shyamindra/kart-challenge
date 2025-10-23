package storage

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"kart-challenge/backend/internal/model"
)

const dbFileName = "kart_challenge.db"

// SQLiteStore implements ProductRepository and OrderRepository using an SQLite database.
type SQLiteStore struct {
	db *sql.DB
	mu sync.RWMutex
}

// NewSQLiteStore creates and initializes a new SQLiteStore.
func NewSQLiteStore(dataDir string) (*SQLiteStore, error) {
	dbPath := filepath.Join(dataDir, dbFileName)

	// Ensure the data directory exists
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create data directory: %w", err)
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	s := &SQLiteStore{db: db}

	if err := s.initSchema(); err != nil {
		return nil, fmt.Errorf("failed to initialize database schema: %w", err)
	}

	// Only initialize products if the table is empty
	if count, _ := s.getProductCount(); count == 0 {
		if err := s.initProducts(); err != nil {
			return nil, fmt.Errorf("failed to initialize products: %w", err)
		}
	}

	return s, nil
}

// initSchema creates the necessary tables if they do not exist.
func (s *SQLiteStore) initSchema() error {
	productTable := `
	CREATE TABLE IF NOT EXISTS products (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		price REAL NOT NULL,
		category TEXT NOT NULL,
		image_json TEXT
	);`

	orderTable := `
	CREATE TABLE IF NOT EXISTS orders (
		id TEXT PRIMARY KEY,
		total REAL NOT NULL,
		discounts REAL NOT NULL,
		items_json TEXT NOT NULL
	);`

	if _, err := s.db.Exec(productTable); err != nil {
		return err
	}
	if _, err := s.db.Exec(orderTable); err != nil {
		return err
	}
	return nil
}

// getProductCount returns the number of products in the database.
func (s *SQLiteStore) getProductCount() (int, error) {
	var count int
	err := s.db.QueryRow("SELECT COUNT(*) FROM products").Scan(&count)
	return count, err
}

// initProducts hardcodes the initial product data into the database.
func (s *SQLiteStore) initProducts() error {
	log.Println("Initializing products in SQLite database...")
	products := []model.Product{
		createProduct(
			"10",
			"Chicken Waffle",
			13.3,
			"Waffle",
			"https://orderfoodonline.deno.dev/public/images/image-waffle-desktop.jpg",
			"https://orderfoodonline.deno.dev/public/images/image-waffle-mobile.jpg",
			"https://orderfoodonline.deno.dev/public/images/image-waffle-thumbnail.jpg",
			"https://orderfoodonline.deno.dev/public/images/image-waffle-tablet.jpg",
		),
		createProduct(
			"11",
			"Beef Burger",
			15.0,
			"Burger",
			"https://orderfoodonline.deno.dev/public/images/image-burger-desktop.jpg",
			"https://orderfoodonline.deno.dev/public/images/image-burger-mobile.jpg",
			"https://orderfoodonline.deno.dev/public/images/image-burger-thumbnail.jpg",
			"https://orderfoodonline.deno.dev/public/images/image-burger-tablet.jpg",
		),
		// Add more products here if needed
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO products(id, name, price, category, image_json) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, p := range products {
		imageJSON, _ := json.Marshal(p.Image)
		_, err = stmt.Exec(*p.Id, *p.Name, *p.Price, *p.Category, string(imageJSON))
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

// createProduct is a helper to create model.Product instances with correct pointer fields.
func createProduct(id, name string, price float32, category, desktop, mobile, thumbnail, tablet string) model.Product {
	sPtr := func(s string) *string { return &s }
	fPtr := func(f float32) *float32 { return &f }

	return model.Product{
		Id:       sPtr(id),
		Name:     sPtr(name),
		Price:    fPtr(price),
		Category: sPtr(category),
		Image: &struct {
			Desktop   *string `json:"desktop,omitempty"`
			Mobile    *string `json:"mobile,omitempty"`
			Tablet    *string `json:"tablet,omitempty"`
			Thumbnail *string `json:"thumbnail,omitempty"`
		}{
			Desktop:   sPtr(desktop),
			Mobile:    sPtr(mobile),
			Thumbnail: sPtr(thumbnail),
			Tablet:    sPtr(tablet),
		},
	}
}

// GetAllProducts returns all available products from the database.
func (s *SQLiteStore) GetAllProducts() ([]model.Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	rows, err := s.db.Query("SELECT id, name, price, category, image_json FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []model.Product{}
	for rows.Next() {
		var p model.Product
		var imageJSON string
		var id, name, category sql.NullString
		var price sql.NullFloat64

		if err := rows.Scan(&id, &name, &price, &category, &imageJSON); err != nil {
			return nil, err
		}

		p.Id = sPtr(id.String)
		p.Name = sPtr(name.String)
		p.Price = fPtr(float32(price.Float64))
		p.Category = sPtr(category.String)

		var image struct {
			Desktop   *string `json:"desktop,omitempty"`
			Mobile    *string `json:"mobile,omitempty"`
			Tablet    *string `json:"tablet,omitempty"`
			Thumbnail *string `json:"thumbnail,omitempty"`
		}
		if err := json.Unmarshal([]byte(imageJSON), &image); err != nil {
			return nil, fmt.Errorf("failed to unmarshal image JSON for product %s: %w", id.String, err)
		}
		p.Image = &image

		products = append(products, p)
	}

	return products, nil
}

// GetProductByID returns a product by its ID from the database.
func (s *SQLiteStore) GetProductByID(id string) (*model.Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	row := s.db.QueryRow("SELECT id, name, price, category, image_json FROM products WHERE id = ?", id)

	var p model.Product
	var imageJSON string
	var name, category sql.NullString
	var price sql.NullFloat64

	if err := row.Scan(&id, &name, &price, &category, &imageJSON); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("product with ID %s not found", id)
		}
		return nil, err
	}

	p.Id = sPtr(id)
	p.Name = sPtr(name.String)
	p.Price = fPtr(float32(price.Float64))
	p.Category = sPtr(category.String)

	var image struct {
		Desktop   *string `json:"desktop,omitempty"`
		Mobile    *string `json:"mobile,omitempty"`
		Tablet    *string `json:"tablet,omitempty"`
		Thumbnail *string `json:"thumbnail,omitempty"`
	}
	if err := json.Unmarshal([]byte(imageJSON), &image); err != nil {
		return nil, fmt.Errorf("failed to unmarshal image JSON for product %s: %w", id, err)
	}
	p.Image = &image

	return &p, nil
}

// CreateOrder adds a new order to the database.
func (s *SQLiteStore) CreateOrder(order model.Order) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	itemsJSON, err := json.Marshal(order.Items)
	if err != nil {
		return fmt.Errorf("failed to marshal order items: %w", err)
	}

	_, err = s.db.Exec(
		"INSERT INTO orders(id, total, discounts, items_json) VALUES(?, ?, ?, ?)",
		*order.Id,
		*order.Total,
		*order.Discounts,
		string(itemsJSON),
	)

	if err != nil {
		return fmt.Errorf("failed to insert order: %w", err)
	}

	return nil
}

// Helper functions for pointers
func sPtr(s string) *string { return &s }
func fPtr(f float32) *float32 { return &f }

// Close closes the database connection.
func (s *SQLiteStore) Close() error {
	return s.db.Close()
}