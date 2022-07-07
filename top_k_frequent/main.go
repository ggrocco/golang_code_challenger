/*
	Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

	Example 1:
	Input: nums = [1,1,1,2,2,3], k = 2
	Output: [1,2]

	Example 2:
	Input: nums = [1], k = 1
	Output: [1]
*/

package main

import (
	"fmt"
	"reflect"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	counters := map[int]int{}
	for _, num := range nums {
		counters[num] = counters[num] + 1
	}

	keys := make([]int, 0, len(counters))
	for key := range counters {
		keys = append(keys, key)
	}

	// buuble
	// for i := len(keys); i > 0; i-- {
	// 	for j := 1; j < i; j++ {
	// 		if counters[keys[j-1]] < counters[keys[j]] {
	// 			intermediate := keys[j]
	// 			keys[j] = keys[j-1]
	// 			keys[j-1] = intermediate
	// 		}
	// 	}
	// }

	sort.SliceStable(keys, func(i, j int) bool {
		return counters[keys[i]] > counters[keys[j]]
	})

	return keys[:k]
}

func main() {
	type test struct {
		nums   []int
		target int
		want   []int
	}

	tests := []test{
		{
			[]int{1, 1, 1, 2, 2, 3},
			2,
			[]int{1, 2},
		},
		{
			[]int{1, 1, 1, 2, 2, 3, 4, 4, 4, 4, 4, 4},
			2,
			[]int{4, 1},
		},
		{
			[]int{1},
			1,
			[]int{1},
		},
	}

	for i, tt := range tests {
		got := topKFrequent(tt.nums, tt.target)
		res := "PASSED"
		if !reflect.DeepEqual(got, tt.want) {
			res = "FAILED"
		}
		fmt.Printf("\n%s case=%d got=%v want=%v", res, i, got, tt.want)
	}
}
