package bfs

import (
	"container/list"
)

func FindOrder(numCourses int, prerequisites [][]int) []int {
	topoOrder := make([]int, 0)
	// 1. 统计每个节点的入度个数，并建立各个节点的映射关系
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	for _, prerequisite := range prerequisites {
		// 理解这两行代码
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
		inDegree[prerequisite[0]]++
	}
	// 2. 将每个入度为0的点放入队列中作为起始节点
	queue := list.New()
	var curCourse int
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue.PushBack(i)
		}
	}
	// 3. 不断从队列中拿出一个点，去掉这个点的所有连边（指向其他点的边），其他点的相应入度 - 1
	for queue.Len() > 0 {
		curCourse = queue.Remove(queue.Front()).(int)
		topoOrder = append(topoOrder, curCourse)
		for _, v := range graph[curCourse] {
			inDegree[v]--
			// 4. 一旦发现新的入度为0的点，丢回队列中
			if inDegree[v] == 0 {
				queue.PushBack(v)
			}
		}
	}
	// 不相等相当于没有修完所有课程，出现有环的情况
	if len(topoOrder) != numCourses {
		return []int{}
	}
	return topoOrder
}
