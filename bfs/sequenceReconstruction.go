package bfs

import (
	"container/list"
)

func sequenceReconstruction(nums []int, sequences [][]int) bool {
	n := len(nums)
	graph := make([][]int, n+1)
	inDegree := make([]int, n+1)
	for _, sequence := range sequences {
		for i := 1; i < len(sequence); i++ {
			x, y := sequence[i-1], sequence[i]
			graph[x] = append(graph[x], y)
			inDegree[y]++
		}
	}
	queue := list.New()
	var curVal int
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			queue.PushBack(i)
		}
	}
	for queue.Len() > 0 {
		if queue.Len() > 1 {
			return false
		}
		curVal = queue.Remove(queue.Front()).(int)
		for _, v := range graph[curVal] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue.PushBack(v)
			}
		}
	}
	return true
}
