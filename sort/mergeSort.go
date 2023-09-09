package sort

// MergeSort 归并排序只能返回新数组，无法原地修改
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]
	return merge(MergeSort(left), MergeSort(right))
}

func merge(arr1, arr2 []int) (mergeArr []int) {
	var index1, index2 int
	for index1 < len(arr1) && index2 < len(arr2) {
		if arr1[index1] <= arr2[index2] {
			mergeArr = append(mergeArr, arr1[index1])
			index1++
		} else {
			mergeArr = append(mergeArr, arr2[index2])
			index2++
		}
	}
	if index1 < len(arr1) {
		mergeArr = append(mergeArr, arr1[index1:]...)
	}
	if index2 < len(arr2) {
		mergeArr = append(mergeArr, arr2[index2:]...)
	}
	return
}
