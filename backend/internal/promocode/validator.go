package promocode

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const (
	minPromoCodeLength = 8
	maxPromoCodeLength = 10
)

// Validator holds the data directory path.
type Validator struct {
	dataDir string
	codes   map[string]int
	mu      sync.RWMutex
}

// NewValidator creates and initializes a new promo code validator.
func NewValidator(dataDir string) *Validator {
	v := &Validator{
		dataDir: dataDir,
		codes:   make(map[string]int),
	}
	// Perform the expensive setup here
	if err := v.setupPromoCodes(); err != nil {
		// Log the error, but don't prevent the server from starting
		// In a real application, you might want to make this a fatal error
		fmt.Printf("Error setting up promo codes: %v\n", err)
	}
	return v
}

// setupPromoCodes performs the one-time expensive setup of downloading, decompressing,
// and building the lookup table for promo codes.
func (v *Validator) setupPromoCodes() error {
	// The files are expected to be in v.dataDir after running prepare_data.sh
	promoCodeFiles := []string{
		"couponbase1",
		"couponbase2",
		"couponbase3",
	}

	// Use a temporary map to store counts before updating the main map
	tempCodes := make(map[string]int)

	for i, filename := range promoCodeFiles {
		filePath := filepath.Join(v.dataDir, filename)
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("failed to open promo code file %s: %w", filePath, err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			code := strings.TrimSpace(scanner.Text())
			if len(code) >= minPromoCodeLength && len(code) <= maxPromoCodeLength {
				// Use bitmask to indicate which file the code was found in
				tempCodes[code] |= (1 << i)
			}
		}

		if scanner.Err() != nil {
			return fmt.Errorf("scanner error for %s: %w", filePath, scanner.Err())
		}
	}

	v.mu.Lock()
	defer v.mu.Unlock()
	v.codes = tempCodes // Replace the map with the fully built one
	return nil
}

// LoadPromoCodes is no longer needed for this strategy, but kept for interface compatibility.
func (v *Validator) LoadPromoCodes() error {
	return nil
}

// IsValid checks if a promo code is valid using the pre-built lookup table.
func (v *Validator) IsValid(code string) bool {
	if len(code) < minPromoCodeLength || len(code) > maxPromoCodeLength {
		return false
	}

	v.mu.RLock()
	defer v.mu.RUnlock()

	bitmask, ok := v.codes[code]
	if !ok {
		return false // Code not found in any file
	}

	// Check if the code was found in at least two files (i.e., at least two bits are set)
	// A bitmask with at least two bits set will have (bitmask & (bitmask - 1)) != 0
	return (bitmask & (bitmask - 1)) != 0
}
