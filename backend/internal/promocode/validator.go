package promocode

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

const (
	minPromoCodeLength = 8
	maxPromoCodeLength = 10
)

// Validator holds the data directory path.
type Validator struct {
	dataDir string
}

// NewValidator creates and initializes a new promo code validator.
func NewValidator(dataDir string) *Validator {
	return &Validator{
		dataDir: dataDir,
	}
}

// LoadPromoCodes is no longer needed for this strategy, but kept for interface compatibility.
func (v *Validator) LoadPromoCodes() error {
	return nil
}

// IsValid checks if a promo code is valid by scanning files on demand.
func (v *Validator) IsValid(code string) bool {
	if len(code) < minPromoCodeLength || len(code) > maxPromoCodeLength {
		return false
	}

	files := []string{
		filepath.Join(v.dataDir, "couponbase1"),
		filepath.Join(v.dataDir, "couponbase2"),
		filepath.Join(v.dataDir, "couponbase3"),
	}

	foundCount := 0
	for _, filePath := range files {
		file, err := os.Open(filePath)
		if err != nil {
			// Log the error, but don't fail the validation entirely
			// log.Printf("Error opening promo code file %s: %v", filePath, err)
			continue
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if strings.TrimSpace(scanner.Text()) == code {
				foundCount++
				break // Found in this file, move to the next
			}
		}

	}

	return foundCount >= 2
}
