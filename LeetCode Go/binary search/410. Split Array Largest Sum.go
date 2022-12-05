package main

func splitArray(nums []int, m int) int {
	// guess answer 猜答案 use binary search
	// slice array into m subarrays
	// [7, 2, 5, 10, 8] k = 2 -> [7,2,5], [10, 8]
	// 切成m份
	sumArr, MaxArr := 0, 0
	for _, num := range nums {
		sumArr += num
		if num > MaxArr {
			MaxArr = num
		}
	}

	//  Now we get max and sum of the array, call binary
	res := binary(nums, m, sumArr, MaxArr)
	return res
}

func binary(nums []int, m int, high int, low int) int {
	res := high
	for low <= high {
		mid := low + (high-low)/2
		if isValid(nums, m, mid) {
			res = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return res
}

func isValid(nums []int, m int, subArraySum int) bool {
	curSum, subArrCount := 0, 1
	for _, num := range nums {
		curSum += num
		if curSum > subArraySum {
			curSum = num
			subArrCount++
			if subArrCount > m {
				return false
			}
		}
	}
	return true
}
