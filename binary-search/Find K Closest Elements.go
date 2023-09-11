package binary_search

func KClosestNumbers(a []int, target int, k int) []int {
	if len(a) < k {
		return nil
	}
	index := getIndex(a, target)
	if index == -1 {
		return nil
	}
	res := make([]int, k)
	left, right := index-1, index
	for i := 0; i < k; i++ {
		if isLeftCloser(a, left, right, target) {
			res[i] = a[left]
			left--
		} else {
			res[i] = a[right]
			right++
		}
	}
	return res
}

func isLeftCloser(a []int, left, right, target int) bool {
	if left < 0 {
		return false
	}
	if right >= len(a) {
		return true
	}
	return target-a[left] <= a[right]-target
}

func getIndex(a []int, target int) int {
	start, end := 0, len(a)-1
	for start+1 < end {
		mid := start + (end-start)/2
		if a[mid] > target {
			end = mid
		} else if a[mid] < target {
			start = mid
		} else {
			start = mid
		}
	}
	return end
}
