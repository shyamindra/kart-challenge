package promocode

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator_IsValid(t *testing.T) {
	// Create dummy data directory and files for testing
	testDataDir := t.TempDir()

	// Create dummy coupon files (uncompressed)
	coupon1Content := []byte("CODEONE123\nCODETWO456\nVALIDCODE1\n")
	coupon2Content := []byte("CODETWO456\nCODEFOUR89\nVALIDCODE1\n")
	coupon3Content := []byte("CODETHREE7\nCODEFIVE01\n")

	err := os.WriteFile(filepath.Join(testDataDir, "couponbase1"), coupon1Content, 0644)
	assert.NoError(t, err)
	err = os.WriteFile(filepath.Join(testDataDir, "couponbase2"), coupon2Content, 0644)
	assert.NoError(t, err)
	err = os.WriteFile(filepath.Join(testDataDir, "couponbase3"), coupon3Content, 0644)
	assert.NoError(t, err)

	validator := NewValidator(testDataDir)

	tests := []struct {
		name     string
		code     string
		expected bool
	}{
		{"Valid code found in 2 files", "CODETWO456", true},
		{"Valid code found in 2 files (another)", "VALIDCODE1", true},
		{"Invalid code found in 1 file", "CODEONE123", false},
		{"Invalid code not found", "NONEXISTENT", false},
		{"Invalid code - too short", "SHORT", false},
		{"Invalid code - too long", "TOOLONGCODE", false},
		{"Invalid code - empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, validator.IsValid(tt.code))
		})
	}
}
