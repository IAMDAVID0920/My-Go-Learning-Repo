package main

func moveZeroes(nums []int) {
	left := 0
	// first put all non-zeros into the array
	for _, val := range nums {
		if val != 0 {
			nums[left] = val
			left++
		}
	}
	// check how many spaces left, fill in with zeros
	for i := left; i < len(nums); i++ {
		nums[i] = 0
	}
}
