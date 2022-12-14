package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	N := len(nums)
	res := nums[0] + nums[1] + nums[N-1]
	sort.Ints(nums)
	fmt.Println(nums)

	for i := 0; i < N-2; i++ {
		left, right := i+1, N-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if abs(sum-target) < abs(res-target) {
				res = sum
			}
			if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}
