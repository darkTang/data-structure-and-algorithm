package dfs

/*
	1. 与 Lowest Common Ancestor 基本一样
	2. 每个节点多了一个指向父节点的指针 Parent
*/

type TreeNodeII struct {
	Val    int
	Left   *TreeNodeII
	Right  *TreeNodeII
	Parent *TreeNodeII
}

func lowestCommonAncestorII(root, p, q *TreeNodeII) *TreeNodeII {
	parentMap := make(map[*TreeNodeII]bool)
	curNode := p
	for curNode != nil {
		parentMap[curNode] = true
		curNode = curNode.Parent
	}
	curNode = q
	for curNode != nil {
		if parentMap[curNode] {
			return curNode
		} else {
			curNode = curNode.Parent
		}
	}
	return nil
}
