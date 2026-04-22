package main

import "errors"

func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("math error: cannot divide by zero")
	}
	return a / b, nil
}
