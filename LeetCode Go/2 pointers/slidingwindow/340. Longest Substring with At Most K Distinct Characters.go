package main

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	tmpMap := make(map[byte]int)
	left, res := 0, 0

	for end := 0; end < len(s); end++ {
		cur := s[end]
		tmpMap[cur]++

		for len(tmpMap) > k {
			leftC := s[left]
			tmpMap[leftC]--
			if tmpMap[leftC] == 0 {
				delete(tmpMap, leftC)
			}
			left++
		}

		res = max(res, end-left+1)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
