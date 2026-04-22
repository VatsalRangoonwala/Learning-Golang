package main

import "fmt"

func main() {
	amount := 1000.15

	discount := 0.0

	if amount >= 1000 {
		discount = 0.20
	} else if amount >= 500 {
		discount = 0.10
	}

	final := amount * (1 - discount)

	fmt.Printf("Final Amount: %.2f\n", final)
}
