package main

import "errors"

func calculateBill(price float64, quantity int, discount float64) (float64, error) {
	if price <= 0 {
		return 0, errors.New("invalid price")
	}
	if quantity <= 0 {
		return 0, errors.New("invalid quantity")
	}
	if discount < 0 {
		return 0, errors.New("invalid discount")
	}

	total := price * float64(quantity)
	final := total * (1 - discount/100)

	return final, nil
}
