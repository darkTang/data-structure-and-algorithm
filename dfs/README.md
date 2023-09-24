# 深度优先搜索（DFS）
## 递归、DFS、回溯、遍历、分治、迭代
- 搜索
  - BFS
  - DFS = 回溯
    - 遍历法
      - 递归
      - 迭代
    - 分治法
      - 递归
      - 迭代

## 二叉树问题的算法
### 考察形态
- 二叉树上求值、求路径
  - [Minimum Subtree](https://www.lintcode.com/problem/596/)
  - [Lowest Common Ancestor](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/)
  - [Lowest Common Ancestor II](https://www.lintcode.com/problem/474/)
  - [Lowest Common Ancestor III](https://www.lintcode.com/problem/578/)

## 二叉树分治法模板
- 分治法通常采用后序遍历
- 分配小弟去做子任务，自己进行结果的汇总
```go
package dfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func divideConquer(root *TreeNode)  {
    if root == nil {
		//处理空树应该返回的结果
    }
	//if root.Left == nil && root.Right == nil {
	//	处理叶子节点应该返回的结果
	//	如果叶子节点返回的结果可以通过两个空节点的返回结果得到就可以省略这一段代码
    //}
	
	//左子树返回结果 = divideConquer(root.Left)
	//右子树返回结果 = divideConquer(root.Right)
    //整棵树的结果 = 按照一定方法合并左右子树的结果
	
	//return 整棵树的结果
}
```