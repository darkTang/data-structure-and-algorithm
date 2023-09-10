package binary_search

func BinarySearch(nums []int, target int) int {
	// write your code here
	if nums == nil || len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1
	for start+1 < end {
		//mid := (start + end) / 2
		mid := start + (end-start)/2 // 相比于上面的写法，当值特别大的时不会出现溢出
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	return -1
}
