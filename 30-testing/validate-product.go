package main

import "errors"

var ErrInvalidPrice = errors.New("invalid price")
var ErrInvalidQuantity = errors.New("invalid quantity")

func ValidateProduct(price float64, quantity int) error {

	if price <= 0 {
		return ErrInvalidPrice
	}

	if quantity <= 0 {
		return ErrInvalidQuantity
	}

	return nil
}
