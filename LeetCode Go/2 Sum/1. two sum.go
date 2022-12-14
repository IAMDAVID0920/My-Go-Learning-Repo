package main

func twoSum(nums []int, target int) []int {
	tmpMap := make(map[int]int)
	for i, num := range nums {
		// When checking specific val in map or not, use this i, ok := map[target-num]
		index, ok := tmpMap[target-num]
		if ok {
			return []int{index, i}
		}
		tmpMap[num] = i
	}
	return []int{}
}
