package main

import "fmt"

func main() {
	productName := "Shirt"
	price := 250.00
	quantity := 2
	discount := 10.00

	total := price * float64(quantity)
	discountAmount := total * discount / 100
	finalPrice := total - discountAmount

	fmt.Printf("Product: %s\nTotal: %.2f\nDiscount: %.0f%%\nFinal Price: %.2f\n", productName, total, discount, finalPrice)
}
