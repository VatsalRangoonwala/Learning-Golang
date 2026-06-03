package main

import (
	"fmt"
)

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

type Order struct {
	ID       string
	Products []Product
	Total    float64
}

func generateOrders() <-chan Order {
	out := make(chan Order)
	orders := []Order{
		{
			Products: []Product{
				{Name: "Wireless Keyboard", Price: 45.99, Quantity: 1}, // Valid
			},
		},
		{
			Products: []Product{
				{Name: "Broken Mouse", Price: -100.00, Quantity: 1}, // INVALID: Negative Price
			},
		},
		{
			Products: []Product{
				{Name: "Office Chair", Price: 150.00, Quantity: 1}, // Valid
			},
		},
		{
			Products: []Product{
				{Name: "Ghost Monitor", Price: 200.00, Quantity: 0}, // INVALID: Zero Quantity
			},
		},
		{
			Products: []Product{
				{Name: "Notebook", Price: 5.50, Quantity: 5}, // Valid
			},
		},
	}

	go func() {
		defer close(out)

		for _, v := range orders {
			out <- v
		}
	}()

	return out
}

func validateOrders(in <-chan Order, errCh chan<- error) <-chan Order {
	out := make(chan Order)

	go func() {
		defer close(out)
		defer close(errCh)

		for v := range in {
			isValid := true

			for _, p := range v.Products {
				if p.Price <= 0 || p.Quantity <= 0 {
					errCh <- fmt.Errorf("invalid Order: %s", p.Name)
					isValid = false
					break
				}
			}
			if isValid {
				out <- v
			}
		}
	}()

	return out
}

func calculateTotal(in <-chan Order) <-chan Order {
	out := make(chan Order)

	go func() {
		defer close(out)

		for v := range in {
			var total float64
			for _, p := range v.Products {
				total += p.Price * float64(p.Quantity)
			}
			v.Total = total
			out <- v
		}
	}()

	return out
}

func assignID(in <-chan Order) <-chan Order {
	out := make(chan Order)

	go func() {
		defer close(out)

		i := 1
		for v := range in {
			v.ID = fmt.Sprintf("ORD %d", i)
			i++
			out <- v
		}
	}()

	return out
}

func main() {
	errCh := make(chan error)
	orders := generateOrders()

	validated := validateOrders(orders, errCh)

	priced := calculateTotal(validated)

	final := assignID(priced)

	for final != nil || errCh != nil {
		select {
		case order, ok := <-final:
			if !ok {
				final = nil
				continue
			}
			fmt.Printf("✅ SUCCESS: %+v\n", order)
		case err, ok := <-errCh:
			if !ok {
				errCh = nil
				continue
			}
			fmt.Printf("❌ ERROR: %v\n", err)
		}
	}
}
