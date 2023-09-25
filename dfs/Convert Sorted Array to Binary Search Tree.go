package dfs

func sortedArrayToBST(nums []int) *TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}
	rootVal := nums[length/2]
	root := &TreeNode{Val: rootVal}

	root.Left = sortedArrayToBST(nums[:length/2])
	root.Right = sortedArrayToBST(nums[length/2+1:])

	return root
}
