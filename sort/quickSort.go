package sort

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	left, right := make([]int, 0), make([]int, 0)
	for i := 1; i < len(arr); i++ {
		if arr[i] >= pivot {
			right = append(right, arr[i])
		} else {
			left = append(left, arr[i])
		}
	}

	return append(QuickSort(left), append([]int{pivot}, QuickSort(right)...)...)
}

/*
	1. 这里的left <= right, arr[left] < pivot，符号要严格区分
	2. 原地修改数组
*/

func QuickSort2(arr []int, start, end int) {
	if start >= end {
		return
	}
	left, right := start, end
	pivot := arr[(start+end)/2]
	for left <= right {
		for left <= right && arr[left] < pivot {
			left++
		}
		for left <= right && arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	QuickSort2(arr, start, right)
	QuickSort2(arr, left, end)
}
