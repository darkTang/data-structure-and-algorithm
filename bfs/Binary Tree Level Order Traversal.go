package bfs

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LevelOrder 单队列实现
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	arr := make([]int, 0)
	st := list.New()
	st.PushBack(root)
	var tempNode *TreeNode
	for st.Len() > 0 {
		length := st.Len()
		for i := 0; i < length; i++ {
			tempNode = st.Remove(st.Front()).(*TreeNode)
			arr = append(arr, tempNode.Val)
			if tempNode.Left != nil {
				st.PushBack(tempNode.Left)
			}
			if tempNode.Right != nil {
				st.PushBack(tempNode.Right)
			}
		}
		res = append(res, arr)
		arr = []int{}
	}
	return res
}

// LevelOrder2 dummyNode实现
func LevelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	arr := make([]int, 0)
	st := list.New()
	st.PushBack(root)
	st.PushBack(nil)
	for st.Len() > 0 {
		// 这边要断言，断言成功就是 *TreeNode 类型，断言失败就是 nil 类型
		// Go 语言中的类型断言不能用于检测一个接口值是否为 nil
		if tempNode, ok := st.Remove(st.Front()).(*TreeNode); ok {
			arr = append(arr, tempNode.Val)
			if tempNode.Left != nil {
				st.PushBack(tempNode.Left)
			}
			if tempNode.Right != nil {
				st.PushBack(tempNode.Right)
			}
		} else {
			res = append(res, arr)
			arr = []int{}
			if st.Len() > 0 { // 做出判断，否则就会死循环
				st.PushBack(nil)
			}
		}
	}
	return res
}
