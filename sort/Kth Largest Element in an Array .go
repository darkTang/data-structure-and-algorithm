package sort

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k)
}

func quickSelect(nums []int, start, end, k int) int {
	if start == end {
		return nums[start]
	}
	left, right := start, end
	pivot := nums[(start+end)/2]
	for left <= right {
		for left <= right && nums[left] > pivot {
			left++
		}
		for left <= right && nums[right] < pivot {
			right--
		}
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	if start+k-1 <= right {
		return quickSelect(nums, start, right, k)
	}
	if start+k-1 >= left {
		return quickSelect(nums, left, end, k-(left-start))
	}
	return nums[right+1]
}
