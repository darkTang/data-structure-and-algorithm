package dfs

import (
	"math"
)

/*
	给 n 个城市(从 1 到 n)，城市和无向道路成本之间的关系为3元组 [A, B, C]（在城市 A 和城市 B 之间有一条路，成本是 C）我们需要从1开始找到的旅行所有城市的付出最小的成本。
	输入:
	n = 3
	tuple = [[1,2,1],[2,3,2],[1,3,3]]
	输出: 3
	说明：最短路是1->2->3
*/

func MinCost(n int, roads [][]int) int {
	graph := constructGraph(n, roads)
	res := 100000
	visited := map[int]bool{1: true}
	var dfs func(city int, visited map[int]bool, cost int)
	dfs = func(city int, visited map[int]bool, cost int) {
		if len(visited) == n {
			res = int(math.Min(float64(res), float64(cost)))
			return
		}
		for i := 1; i < len(graph[city]); i++ {
			if !visited[i] {
				visited[i] = true
				//将整数的最大值加1时，会发生溢出。溢出是指计算结果超出了变量所能表示的范围。在Go语言中，溢出是一种运行时错误，因为它可能导致未定义的行为。
				dfs(i, visited, cost+graph[city][i])
				delete(visited, i)
			}
		}
	}
	dfs(1, visited, 0)
	return res
}

// 带权无向图  邻接矩阵法表示
func constructGraph(n int, roads [][]int) [][]int {
	graph := make([][]int, n+1)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]int, n+1)
		for j := 0; j < len(graph[0]); j++ {
			graph[i][j] = 100000
		}
	}
	for _, road := range roads {
		i, j, cost := road[0], road[1], road[2]
		graph[i][j] = int(math.Min(float64(graph[i][j]), float64(cost)))
		graph[j][i] = int(math.Min(float64(graph[j][i]), float64(cost)))
	}
	return graph
}
