package binary_search

import "math"

// MountainSequence 不能存在相等的值
func MountainSequence(nums []int) int {
	start, end := 0, len(nums)-1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] > nums[mid+1] {
			end = mid
		} else {
			start = mid
		}
	}
	return int(math.Max(float64(nums[start]), float64(nums[end])))
}
