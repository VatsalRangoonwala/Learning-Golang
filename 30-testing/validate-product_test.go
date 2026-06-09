package main

import (
	"errors"
	"testing"
)

func TestValidateProduct(t *testing.T) {
	tests := []struct {
		name     string
		price    float64
		quantity int
		expected error
	}{
		{
			name:     "Valid product",
			price:    100.0,
			quantity: 1,
			expected: nil,
		},
		{
			name:     "Invalid: Zero price",
			price:    0.0,
			quantity: 1,
			expected: ErrInvalidPrice,
		},
		{
			name:     "Invalid: Negative price",
			price:    -50.0,
			quantity: 1,
			expected: ErrInvalidPrice,
		},
		{
			name:     "Invalid: Zero quantity",
			price:    100.0,
			quantity: 0,
			expected: ErrInvalidQuantity,
		},
		{
			name:     "Invalid: Negative quantity",
			price:    100.0,
			quantity: -5,
			expected: ErrInvalidQuantity,
		},
	}

	for _, tt := range tests {
		
		t.Run(tt.name, func(t *testing.T) {
			
			result := ValidateProduct(tt.price, tt.quantity)

			if !errors.Is(result, tt.expected) {
				t.Errorf("\nInputs  : price=%.2f, quantity=%d\nExpected: %v\nGot     : %v",
					tt.price, tt.quantity, tt.expected, result)
			}
		})
	}
}