package dfs

import "container/list"

func kthSmallest(root *TreeNode, k int) int {
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
	return nums[k-1]
}

func kthSmallest2(root *TreeNode, k int) int {
	st := list.New()
	curNode := root
	tmp := 0
	for st.Len() > 0 || curNode != nil {
		if curNode != nil {
			st.PushBack(curNode)
			curNode = curNode.Left
		} else {
			curNode = st.Remove(st.Back()).(*TreeNode)
			tmp++
			if tmp == k {
				return curNode.Val
			}
			curNode = curNode.Right
		}
	}
	return 0
}
