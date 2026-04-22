package main

import "fmt"

func main() {
	nums := []int{10, 45, 23, 89, 12}
	if len(nums) == 0 {
		fmt.Println("Empty slice")
		return
	}
	largest := nums[0]

	for _, val := range nums {
		if val > largest {
			largest = val
		}
	}
	fmt.Println("Largest:", largest)
}
