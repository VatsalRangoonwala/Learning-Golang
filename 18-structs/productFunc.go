package main

import "errors"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func getTotal(p Product) (float64, error) {
	if p.Price <= 0 {
		return 0, errors.New("Invalid Price")
	}
	if p.Quantity <= 0 {
		return 0, errors.New("Invalid Quantity")
	}
	return p.Price * float64(p.Quantity), nil
}
