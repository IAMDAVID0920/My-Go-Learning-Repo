package main

func maxArea(height []int) int {
	// 2 pointer brute force?
	size := len(height)
	left, right := 0, size-1
	res := 0
	for left < right {
		// compare min to determin the edge
		minH := compareMin(height[left], height[right])
		width := right - left
		res = compareMax(res, minH*width)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return res

}

func compareMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func compareMin(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
