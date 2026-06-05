package main

import "testing"

func TestMultiply(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{2, 2, 4},
		{10, 3, 30},
		{7, 8, 56},
		{0, 5, 0},
		{-2, 3, -6},
	}

	for _, test := range tests {
		result := Multiply(test.a, test.b)

		if result != test.expected {
			t.Errorf("expected %d, got %d", test.expected, result)
		}
	}
}
