package dfs

func permute(nums []int) [][]int {
	var res [][]int
	var dfs func(path []int, index int)
	visited := make([]bool, len(nums)) // 标记元素，避免在for循环中重复选择
	dfs = func(path []int, index int) {
		if index == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		for i := 0; i < len(nums); i++ {
			if !visited[i] {
				path = append(path, nums[i])
				visited[i] = true
				dfs(path, index+1)
				path = path[:len(path)-1]
				visited[i] = false
			}
		}
	}
	dfs([]int{}, 0)
	return res
}
