/*
	Given an array of integers nums which is sorted in ascending order, and an integer target,
	write a function to search target in nums. If target exists, then return its index.
	Otherwise, return -1.

	Example 1:
	Input: nums = [-1,0,3,5,9,12], target = 9
	Output: 4
	Explanation: 9 exists in nums and its index is 4

	Example 2:
	Input: nums = [-1,0,3,5,9,12], target = 2
	Output: -1
	Explanation: 2 does not exist in nums so return -1
*/

package main

import (
	"fmt"
)

// func search(nums []int, target int) int {
// 	for i, k := range nums {
// 		if k == target {
// 			return i
// 		}
// 	}
// 	return -1
// }

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	type test struct {
		nums   []int
		target int
		want   int
	}

	tests := []test{
		{
			[]int{-1, 0, 3, 5, 9, 12},
			9,
			4,
		},
		{
			[]int{-1, 0, 3, 5, 9, 12},
			2,
			-1,
		},
		{
			[]int{-1, 0, 3, 5, 9, 12},
			3,
			2,
		},
		{
			[]int{-1, 0, 3, 5, 9, 12},
			5,
			3,
		},
		{
			[]int{-1, 0, 3, 9, 12},
			12,
			4,
		},
		{
			[]int{-1, 0, 3, 9, 12},
			-1,
			0,
		},
	}

	for i, tt := range tests {
		got := search(tt.nums, tt.target)
		res := "PASSED"
		if got != tt.want {
			res = "FAILED"
		}
		fmt.Printf("\n%s case=%d got=%v want=%v", res, i, got, tt.want)
	}
}
