package main

import "fmt"

func lengthOfLIS(nums []int) int {
	// 2pointer dp array
	//         l   r
	// input: [10, 9, 2, 5, 3, 7, 101, 18]
	// dp:    [1,  1, 1, 1, 1, 1,  1,  1]
	// dp:    [1,  1, 1, 2, 2, 3,  3,  4]

	size := len(nums)
	if size == 0 {
		return 0
	}
	// fill in dp array with value 1
	dp := make([]int, len(nums))
	fmt.Println(dp)
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	fmt.Println(dp)
	res := 1

	// 2 pointer
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
				res = max(res, dp[i])
			}
		}
	}
	fmt.Println(dp)
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
