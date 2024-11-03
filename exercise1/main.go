// Given an integer array, nums, move all of the zeros to the end of it
// while maintaing the relative order of non-zero elements

// Plan
// 1. Go through the array and fill all non zero items
// 2. Fill with zeros second partition

package main

import (
	"fmt"
)

func moveZerosToEnd(arr []int) {
	fill_index := 0

	// Go thourgh first partion and set non-zero elements
	for _, value := range arr {
		if value != 0 {
			arr[fill_index] = value
			fill_index++
		}
	}

	// Set zero elements from last non-zero till the end
	for i := fill_index; i < len(arr); i++ {
		arr[i] = 0
	}
}

func main() {
	var arr = []int{0, 1, 2, 3, 40, 0, 2, 3, 0}

	fmt.Println(arr)
	moveZerosToEnd(arr)
	fmt.Println(arr)
}
