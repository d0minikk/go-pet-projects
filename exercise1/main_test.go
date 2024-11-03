package main

import (
	"testing"
	"reflect"
)

func TestMoveZerosToEndWithZerosOnLeft(t *testing.T) {
	var nums = []int{0, 0, 0, 0, 40, 1, 2, 3, 10}
	moveZerosToEnd(nums)

	var correctTransformedNums = []int{40, 1, 2, 3, 10, 0, 0, 0, 0}

	if !reflect.DeepEqual(nums, correctTransformedNums) {
		t.Errorf("Numbers array is not transformed correctly %v", nums)
	}
}

func TestMoveZerosToEndWithZerosOnRight(t *testing.T) {
	var nums = []int{40, 1, 2, 3, 10, 0, 0, 0, 0}
	moveZerosToEnd(nums)

	var correctTransformedNums = []int{40, 1, 2, 3, 10, 0, 0, 0, 0}

	if !reflect.DeepEqual(nums, correctTransformedNums) {
		t.Errorf("Numbers array is not transformed correctly %v", nums)
	}
}

func TestMoveZerosToEndWithNoZeros(t *testing.T) {
	var nums = []int{40, 1, 2, 3, 10}
	moveZerosToEnd(nums)

	var correctTransformedNums = []int{40, 1, 2, 3, 10}

	if !reflect.DeepEqual(nums, correctTransformedNums) {
		t.Errorf("Numbers array is not transformed correctly %v", nums)
	}
}

func TestMoveZerosToEndWithOnlyZeros(t *testing.T) {
	var nums = []int{0, 0, 0}
	moveZerosToEnd(nums)

	var correctTransformedNums = []int{0, 0, 0}

	if !reflect.DeepEqual(nums, correctTransformedNums) {
		t.Errorf("Numbers array is not transformed correctly %v", nums)
	}
}
