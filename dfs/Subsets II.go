package dfs

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var dfs func(subsets []int, startIndex int)
	dfs = func(subsets []int, startIndex int) {
		tmp := make([]int, len(subsets))
		copy(tmp, subsets)
		res = append(res, tmp)
		for i := startIndex; i < len(nums); i++ {
			if i != 0 && nums[i] == nums[i-1] && i > startIndex {
				continue
			}
			subsets = append(subsets, nums[i])
			dfs(subsets, i+1)
			subsets = subsets[:len(subsets)-1]
		}
	}
	dfs([]int{}, 0)
	return res
}
