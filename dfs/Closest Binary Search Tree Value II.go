package dfs

// 给定一棵非空二叉搜索树以及一个target值，找到BST中最接近给定值的k个数
// 给出的目标值为浮点数
// k总是合理的，k<=总节点数
// 我们可以保证只有唯一一个最接近给定值的k个值的集合

/*
	输入：root = {3,1,4,#,2} and target = 0.275 and k = 2
	输出：[1,2]
*/

func closestKValue(root *TreeNode, target float64, k int) []int {
	var inOrder func(root *TreeNode)
	nums := make([]int, 0)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Left)
		nums = append(nums, root.Val)
		inOrder(root.Right)
	}
	inOrder(root)

	index := getIndex(nums, target)
	if index == -1 {
		return nil
	}
	res := make([]int, k)
	left, right := index-1, index
	for i := 0; i < k; i++ {
		if isLeftCloser(nums, left, right, target) {
			res[i] = nums[left]
			left--
		} else {
			res[i] = nums[right]
			right++
		}
	}
	return res
}

func isLeftCloser(nums []int, left, right int, target float64) bool {
	if left < 0 {
		return false
	}
	if right > len(nums) {
		return true
	}
	return target-float64(nums[left]) <= float64(nums[left])-target
}

func getIndex(nums []int, target float64) int {
	start, end := 0, len(nums)-1
	for start+1 < end {
		mid := start + (end-start)/2
		if float64(nums[mid]) > target {
			end = mid
		} else if float64(nums[mid]) < target {
			start = mid
		} else {
			start = mid
		}
	}
	return end
}
