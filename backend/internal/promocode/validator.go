package promocode

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"net/http"
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
	promoCodeURLs := []string{
		"https://orderfoodonline-files.s3.ap-southeast-2.amazonaws.com/couponbase1.gz",
		"https://orderfoodonline-files.s3.ap-southeast-2.amazonaws.com/couponbase2.gz",
		"https://orderfoodonline-files.s3.ap-southeast-2.amazonaws.com/couponbase3.gz",
	}

	// Use a temporary map to store counts before updating the main map
	// This avoids locking the main map for each file processing
	tempCodes := make(map[string]int)

	for i, url := range promoCodeURLs {
		resp, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("failed to download %s: %w", url, err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to download %s: status code %d", url, resp.StatusCode)
		}

		gzReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to create gzip reader for %s: %w", url, err)
		}
		defer gzReader.Close()

		scanner := bufio.NewScanner(gzReader)
		for scanner.Scan() {
			code := strings.TrimSpace(scanner.Text())
			if len(code) >= minPromoCodeLength && len(code) <= maxPromoCodeLength {
				// Use bitmask to indicate which file the code was found in
				tempCodes[code] |= (1 << i)
			}
		}

		if scanner.Err() != nil {
			return fmt.Errorf("scanner error for %s: %w", url, scanner.Err())
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
