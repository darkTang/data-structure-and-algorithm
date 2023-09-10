package binary_search

func lastPosition(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1
	for start+1 < end {
		mid := (start + end) / 2
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			start = mid // ****
		}
	}
	if nums[end] == target {
		return end
	}
	if nums[start] == target {
		return start
	}
	return -1
}
