package dfs

import "math"

func isBalanced(root *TreeNode) bool {
	if getDepth(root) == -1 {
		return false
	} else {
		return true
	}
}

func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftDepth := getDepth(node.Left)
	rightDepth := getDepth(node.Right)
	if leftDepth == -1 || rightDepth == -1 {
		return -1
	}
	if math.Abs(float64(leftDepth)-float64(rightDepth)) > 1 {
		return -1
	}
	maxDepth := int(math.Max(float64(leftDepth), float64(rightDepth)))
	return maxDepth + 1
}
