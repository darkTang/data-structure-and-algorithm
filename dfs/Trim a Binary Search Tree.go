package dfs

// 定义 递归函数就是满足 [low, high] 返回的子树
// 第一步：明确递归函数含义
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	leftNode := trimBST(root.Left, low, high)
	rightNode := trimBST(root.Right, low, high)

	if root.Val >= low && root.Val <= high {
		root.Left = leftNode
		root.Right = rightNode
		return root
	}
	if root.Val < low {
		return rightNode
	}
	if root.Val > high {
		return leftNode
	}
	return nil
}
