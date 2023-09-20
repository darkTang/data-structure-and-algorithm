package bfs

import "container/list"

func NumIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 || grid[0] == nil || len(grid[0]) == 0 {
		return 0
	}
	directs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	visited := make(map[[2]int]bool)
	xLen := len(grid)
	yLen := len(grid[0])
	isLands := 0
	for i := 0; i < xLen; i++ {
		for j := 0; j < yLen; j++ {
			if grid[i][j] == '1' && !visited[[2]int{i, j}] {
				bfs(grid, i, j, visited, directs)
				isLands++
			}
		}
	}
	return isLands
}

func bfs(grid [][]byte, x, y int, visited map[[2]int]bool, directs [][]int) {
	queue := list.New()
	queue.PushBack([]int{x, y})
	visited[[2]int{x, y}] = true
	var curCoordinate []int
	for queue.Len() > 0 {
		curCoordinate = queue.Remove(queue.Front()).([]int)
		// 遍历上下左右四个方向
		for _, direct := range directs {
			newX := curCoordinate[0] + direct[0]
			newY := curCoordinate[1] + direct[1]
			if !visited[[2]int{newX, newY}] && isValid(grid, newX, newY) {
				queue.PushBack([]int{newX, newY})
				visited[[2]int{newX, newY}] = true
			}
		}
	}
}

func isValid(grid [][]byte, x, y int) bool {
	n := len(grid)
	m := len(grid[0])
	if x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == '0' {
		return false
	}
	return true
}
