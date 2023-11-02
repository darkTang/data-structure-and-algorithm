package dfs

import "sort"

// 通过判断每一个数是否放入的状态
func subsets(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var dfs func(list []int, index int)
	dfs = func(list []int, index int) {
		if index == len(nums) {
			tmp := make([]int, len(list))
			copy(tmp, list)
			res = append(res, tmp)
			return
		}
		list = append(list, nums[index])
		dfs(list, index+1)
		list = list[:len(list)-1]
		dfs(list, index+1)
	}
	dfs([]int{}, 0)
	return res
}

// 逐个放入数的思路 通用方法
func subsetsII(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var dfs func(list []int, startIndex int) // startIndex 避免重复取元素，只能往后取，不能往前取
	dfs = func(list []int, startIndex int) {
		tmp := make([]int, len(list))
		copy(tmp, list)
		res = append(res, tmp)
		for i := startIndex; i < len(nums); i++ {
			list = append(list, nums[i])
			dfs(list, i+1) // 这里是 i+1，并不是 startIndex+1
			list = list[:len(list)-1]
		}
	}
	dfs([]int{}, 0)
	return res
}
