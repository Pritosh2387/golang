package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, found := numMap[complement]; found {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return nil 
}

func main() {
	// Example test case
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Printf("Indices: %v\n", result) // Expected: [0 1]
}