package binary_search

import "math"

func FindMin(nums []int) int {
	start, end := 0, len(nums)-1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] > nums[end] {
			start = mid
		} else {
			end = mid
		}
	}
	return int(math.Min(float64(nums[start]), float64(nums[end])))
}
