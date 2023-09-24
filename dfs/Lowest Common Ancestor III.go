package dfs

/*
	1. 与 Lowest Common Ancestor 基本一样
	2. 这两个节点未必都存在
*/

func lowestCommonAncestorIII(root, p, q *TreeNode) *TreeNode {
	pExist, qExist, node := helper(root, p, q)
	if pExist && qExist {
		return node
	}
	return nil
}

func helper(root, p, q *TreeNode) (bool, bool, *TreeNode) {
	if root == nil {
		return false, false, nil
	}

	leftPExist, leftQExist, left := helper(root.Left, p, q)
	rightPExist, rightQExist, right := helper(root.Right, p, q)

	pExist := leftPExist || rightPExist || root == p
	qExist := leftQExist || rightQExist || root == q

	if root == p || root == q {
		return pExist, qExist, root
	}

	if left != nil && right != nil {
		return pExist, qExist, root
	}
	if left != nil {
		return pExist, qExist, left
	}
	if right != nil {
		return pExist, qExist, right
	}
	return pExist, qExist, nil
}
