package main

func searchInsert(nums []int, target int) int {
	// onlogn means we need to use binary search

	left, mid, right := 0, 0, len(nums)-1

	for left <= right {
		mid = left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if nums[mid] > target {
		// nums[mid] == 3, target = 2, insert 2 at nums[mid]'s pos
		return mid
	} else {
		// nums[mid] <= target 4/5 <-> 5
		return mid + 1
	}

}
