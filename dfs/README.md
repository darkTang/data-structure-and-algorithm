# 二叉树分治法模板
## 分治法通常采用后序遍历
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