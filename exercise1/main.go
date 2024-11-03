// Given an integer array, nums, move all of the zeros to the end of it
// while maintaing the relative order of non-zero elements

// Plan
// 1. Go through the array and fill all non zero items
// 2. Fill with zeros second partition

package main

import (
	"fmt"
)

func main() {
	var arr = []int{0, 1, 2, 3, 40, 0, 2, 3, 0}

	fmt.Println(arr)

	fill_index := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[fill_index] = arr[i]
			fill_index++
		}
	}

	for i := fill_index; i < len(arr); i++ {
		arr[i] = 0
	}

	fmt.Println(arr)
}

