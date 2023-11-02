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
  - [Lowest Common Ancestor](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/)
  - [Lowest Common Ancestor II](https://www.lintcode.com/problem/474/)
  - [Lowest Common Ancestor III](https://www.lintcode.com/problem/578/)
- 二叉树结构变化
  - [Flatten Binary Tree to Linked List](https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/)
- 二叉搜索树（BST）：左子树所有节点的值都要小于根节点的值；右子树所有节点的值都要大于根节点的值。左右子树又各是一棵二叉搜索树。如果出现相同的值，既可以放在左子树，也可以放在右子树。(BST尽量不要出现相同的值)
  - 题目
    - Build - [Convert Sorted Array to Binary Search Tree](https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/)
    - Insert - [Insert into a Binary Search Tree](https://leetcode.cn/problems/insert-into-a-binary-search-tree/)
    - Search - [Search in a Binary Search Tree](https://leetcode.cn/problems/search-in-a-binary-search-tree/)
    - Delete - [Trim a Binary Search Tree](https://leetcode.cn/problems/trim-a-binary-search-tree/)
    - Iterate - [Binary Search Tree Iterator](https://leetcode.cn/problems/binary-search-tree-iterator/)
    - [Kth Smallest Element in a BST](https://leetcode.cn/problems/kth-smallest-element-in-a-bst/)
    - [Closest Binary Search Tree Value](https://leetcode.cn/problems/closest-binary-search-tree-value/)
    - [Closest Binary Search Tree Value II](https://leetcode.cn/problems/closest-binary-search-tree-value-ii/)
- 平衡二叉树（AVL）：
  - 前提它是一棵二叉搜索树
  - 它是一颗空树或它的左右两个子树的高度差的绝对值不超过1
  - 左右两个子树都是一棵平衡二叉树
  
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

## 组合类DFS
- 子集问题
  - 题目
    - [Subsets](https://leetcode.cn/problems/subsets/)
    - [Subsets II](https://leetcode.cn/problems/subsets-ii/)

## 排列类DFS
- 全排列问题
  - 题目
    - [Permutations](https://leetcode.cn/problems/permutations/description/)