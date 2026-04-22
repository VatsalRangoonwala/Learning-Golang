package main

import "fmt"

func calculate(a, b int) (int, int, int, float64) {
	if b == 0 {
		return a + b, a - b, a * b, 0
	}
	return a + b, a - b, a * b, float64(a) / float64(b)
}

func main() {
	nums := []int{10, 45, 23, 89, 12}

	fmt.Println(calculate(10, 5))
	fmt.Println(isEven(10))
	fmt.Println(findMax(nums))
	fmt.Println(safeDivide(10, 2))
	fmt.Println(calculateBill(149.50, 10, 20))
}
