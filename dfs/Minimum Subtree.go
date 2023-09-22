package dfs

import "math"

/*
	给你一棵二叉树，找到和为最小的子树，返回其根节点（不是根节点的和）
	保证只有一棵和最小的子树
	并且给出的二叉树不是一棵空树
*/

// 全局变量版本
func findSubtree(root *TreeNode) *TreeNode {
	minSum := math.MaxInt
	var minRoot *TreeNode
	var getTreeSum func(node *TreeNode) int
	getTreeSum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftSum := getTreeSum(node.Left)
		rightSum := getTreeSum(node.Right)
		rootSum := leftSum + rightSum + node.Val
		if rootSum < minSum {
			minSum = rootSum
			minRoot = node
		}
		return rootSum
	}
	getTreeSum(root)
	return minRoot
}

func findSubtree2(root *TreeNode) *TreeNode {
	_, _, subtree := getTreeSum(root)
	return subtree
}

func getTreeSum(node *TreeNode) (int, int, *TreeNode) {
	if node == nil {
		return math.MaxInt, 0, nil
	}
	leftMinSum, leftSum, leftSubtree := getTreeSum(node.Left)
	rightMinSum, rightSum, rightSubtree := getTreeSum(node.Right)
	rootSum := leftSum + rightSum + node.Val
	// 如果左子树最小，返回左子树
	if leftMinSum <= rootSum {
		return leftMinSum, rootSum, leftSubtree
	}
	// 如果右子树最小，返回右子树
	if rightMinSum <= rootSum {
		return rightMinSum, rootSum, rightSubtree
	}
	// 如果当前树最小，返回当前树
	return rootSum, rootSum, node
}
