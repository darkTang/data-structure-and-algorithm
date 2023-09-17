package dfs

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BinaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	if root == nil {
		return res
	}
	findPaths(root, "", &res)
	return res
}

func findPaths(node *TreeNode, s string, res *[]string) { // 必须接收的是一个切片
	if node.Left == nil && node.Right == nil {
		v := s + strconv.Itoa(node.Val)
		*res = append(*res, v)
		return
	}
	s = s + strconv.Itoa(node.Val) + "->"
	if node.Left != nil {
		findPaths(node.Left, s, res)
	}
	if node.Right != nil {
		findPaths(node.Right, s, res)
	}
}

// BinaryTreePaths2 分治法思想，采用后序遍历
func BinaryTreePaths2(root *TreeNode) []string {
	paths := make([]string, 0)
	if root == nil {
		return paths
	}
	if root.Left == nil && root.Right == nil {
		paths = append(paths, ""+strconv.Itoa(root.Val))
		return paths
	}
	for _, leftPath := range BinaryTreePaths2(root.Left) {
		paths = append(paths, strconv.Itoa(root.Val)+"->"+leftPath)
	}
	for _, rightPath := range BinaryTreePaths2(root.Right) {
		paths = append(paths, strconv.Itoa(root.Val)+"->"+rightPath)
	}
	return paths
}
