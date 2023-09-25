package dfs

func flatten(root *TreeNode) {
	dfs(root)
}

func dfs(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	leftFlattenNode := dfs(root.Left)
	rightFlattenNode := dfs(root.Right)

	if root.Left != nil {
		leftFlattenNode.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}

	if rightFlattenNode != nil {
		return rightFlattenNode
	}
	if leftFlattenNode != nil {
		return leftFlattenNode
	}
	return root
}
