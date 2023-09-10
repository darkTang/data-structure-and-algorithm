package binary_search

func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, start, end, target int) int {
	if start > end {
		return -1
	}

	mid := (start + end) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[mid] > target {
		return binarySearch(nums, start, mid-1, target)
	}
	return binarySearch(nums, mid+1, end, target)
}

func search2(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func search3(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start < end {
		mid := (start + end) / 2
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			return mid
		}
	}
	if nums[start] == target { // 当数组只有一个元素时
		return start
	}
	return -1

}
