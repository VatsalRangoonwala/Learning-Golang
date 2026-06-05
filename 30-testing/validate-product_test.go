package main

import (
	"testing"
)

func TestValidateProduct(t *testing.T) {
	tests := []struct {
		a        float64
		b        int
		expected error
	}{
		{100, 1, nil},
		{0, 1, ErrInvalidPrice},
		{-50, 1, ErrInvalidPrice},
		{100, 0, ErrInvalidQuantity},
		{100, -5, ErrInvalidQuantity},
	}

	for _, test := range tests {
		result := ValidateProduct(test.a, test.b)

		if result != test.expected {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}
