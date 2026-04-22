package main

import (
	"errors"
)

func findMax(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("Empty slice")
	}
	max := nums[0]

	for _, val := range nums {
		if val > max {
			max = val
		}
	}
	return max, nil
}
