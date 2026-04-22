package main

import "fmt"

func main() {
	productName := "shirt"
	price := 149.50
	quantity := 5
	total := price * float64(quantity)

	fmt.Printf("Your total bill of %s is %.2f\n", productName, total)
}
