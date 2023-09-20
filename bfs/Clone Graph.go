package bfs

import "container/list"

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	queue := list.New()
	// 一旦入队，需要立马标记
	queue.PushBack(node)
	visited := make(map[*Node]*Node)
	visited[node] = &Node{
		Val:       node.Val,
		Neighbors: []*Node{},
	}

	var curNode *Node

	for queue.Len() > 0 {
		// 这里再加一个for循环，可以进行分层
		curNode = queue.Remove(queue.Front()).(*Node)
		for _, neighbor := range curNode.Neighbors {
			if _, ok := visited[neighbor]; !ok {
				//一旦入队，需要立马标记
				queue.PushBack(neighbor)
				visited[neighbor] = &Node{
					Val:       neighbor.Val,
					Neighbors: []*Node{},
				}
			}
			visited[curNode].Neighbors = append(visited[curNode].Neighbors, visited[neighbor])
		}
	}
	return visited[node]
}
