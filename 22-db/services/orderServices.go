package services

import (
	"context"
	"errors"
	"project/models"
	"project/repository"
	"time"

	"github.com/google/uuid"
)

func totalPrice(products []models.Product) float64 {
	var total float64
	for _, p := range products {
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

	order.ID = uuid.New().String()
	order.CreatedAt = now
	order.Total = totalPrice(order.Products)
	order.Status = "pending"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := repository.InsertOrder(ctx, order)
	if err != nil {
		return order, err
	}

	return order, nil
}

func GetOrders() ([]models.Order, error) {
	return repository.GetAllOrders()
}

func GetOrder(id string) (models.Order, error) {
	return repository.GetOrderByID(id)
}