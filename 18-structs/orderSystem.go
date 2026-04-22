package main

import "fmt"

type Order struct {
	Product Product
	User    User
}

func (o Order) summary() (string, error) {
	total, err := getTotal(o.Product)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s ordered %s worth %.2f", o.User.Name, o.Product.Name, total), nil
}
