package main

import "fmt"

func main() {
	nums := []int{2, 5, 8, 11, 14}
	if len(nums) == 0 {
		fmt.Println("Empty slice")
		return
	}

	for _, val := range nums {
		if val%2 == 0 {
			fmt.Println(val)
		}
	}
}
