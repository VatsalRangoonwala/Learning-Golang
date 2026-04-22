package main

import "errors"

func (p Product) finalPrice(discount float64) (float64, error) {
	if discount < 0 {
		return 0, errors.New("Invalid Discount")
	}
	total, err := getTotal(p)
	if err != nil {
		return 0, err
	}
	return total * (1 - discount/100), nil
}
