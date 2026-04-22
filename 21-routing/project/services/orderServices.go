package services

import (
	"errors"
	"fmt"
	"project/models"
	"time"
)

func totalPrice(o models.Order) float64 {
	var total float64
	for _, p := range o.Products {
		total += p.Price * float64(p.Quantity)
	}
	return total
}

func ProcessOrder(order models.Order) (models.Order, error) {

	if len(order.Products) == 0 {
		return order, errors.New("no products provided")
	}

	for _, p := range order.Products {
		if p.Name == "" || p.Price <= 0 || p.Quantity <= 0 {
			return order, errors.New("Invalid product data")
		}
	}

	if order.User.Name == "" || order.User.Email == "" {
		return order, errors.New("Invalid User data")
	}

	now := time.Now()

	order.ID = fmt.Sprintf("ORD-%d", now.UnixNano())
	order.CreatedAt = now
	order.Total = totalPrice(order)

	return order, nil
}
