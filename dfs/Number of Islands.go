package dfs

func NumIslands(grid [][]byte) int {
	num := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				getOtherIslands(i, j, grid)
				num++
			}
		}
	}
	return num
}

func getOtherIslands(i, j int, grid [][]byte) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return
	}
	if grid[i][j] != '1' {
		return
	}
	grid[i][j] = '2' // 标记为2，说明该点已经遍历过了，但是缺点就是它修改了原数据，尤其是在工程问题上
	getOtherIslands(i+1, j, grid)
	getOtherIslands(i-1, j, grid)
	getOtherIslands(i, j+1, grid)
	getOtherIslands(i, j-1, grid)
}
