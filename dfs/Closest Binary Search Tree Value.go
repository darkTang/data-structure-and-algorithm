package dfs

import "math"

// 给定一棵非空二叉搜索树以及一个target值，找到BST中最接近给定值的节点值
// 给出的目标值为浮点数
// 我们可以保证只有唯一一个最接近给定值的节点

/*
	输入：root = {5,4,9,2,#,8,10} and target = 6.124780
	输出：5
*/

func ClosestValue(root *TreeNode, target float64) int {
	if root == nil {
		return 0
	}
	upper, lower := root, root
	for root != nil {
		if float64(root.Val) > target {
			upper = root
			root = root.Left
		} else if float64(root.Val) < target {
			lower = root
			root = root.Right
		} else {
			return root.Val
		}
	}
	if math.Abs(float64(upper.Val)-target) > math.Abs(target-float64(lower.Val)) {
		return lower.Val
	}
	return upper.Val
}

/*
	特别简单好理解的方法，非递归： 如果当前root值比target大，就暂且把这个root值当成上限，
	然后往左边走 如果当前root值比target小，就暂且把这个root值当成下限，然后往右边走 左右摇摆着走，
	知道发现两个最接近target的值
*/
