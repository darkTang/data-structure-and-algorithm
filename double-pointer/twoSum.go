package double_pointer

import "sort"

// hashMap
func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, num := range nums {
		if index, ok := hashMap[num]; ok {
			return []int{index, i}
		}
		hashMap[target-num] = i
	}
	return []int{-1, -1}
}

// sort + 相向双指针
func twoSum2(nums []int, target int) []int {
	arrs := make([][]int, 0)
	for i, num := range nums {
		temp := []int{num, i}
		arrs = append(arrs, temp)
	}
	sort.Slice(arrs, func(i, j int) bool {
		return arrs[i][0] < arrs[j][0]
	})
	var left, right = 0, len(arrs) - 1
	for left < right {
		if arrs[left][0]+arrs[right][0] > target {
			right--
		} else if arrs[left][0]+arrs[right][0] < target {
			left++
		} else {
			return []int{arrs[left][1], arrs[right][1]}
		}
	}
	return []int{-1, -1}
}
