package dfs

import "container/list"

// 迭代法
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	curNode := root
	st := list.New() // 栈的思想
	for curNode != nil || st.Len() > 0 {
		if curNode != nil {
			st.PushBack(curNode)
			curNode = curNode.Left
		} else {
			curNode = st.Remove(st.Back()).(*TreeNode)
			res = append(res, curNode.Val)
			curNode = curNode.Right
		}
	}
	return res
}
