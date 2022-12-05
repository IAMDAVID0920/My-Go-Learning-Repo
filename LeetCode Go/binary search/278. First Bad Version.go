package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool {
	// pseudo
	return false
}

func firstBadVersion(n int) int {
	left, right := 1, n

	for left <= right {
		mid := left + (right-left)/2

		if !isBadVersion(mid) {
			// means that we still not finding any bad versions, find it right-half
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
