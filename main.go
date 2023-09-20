package main

import (
	"data-structure-and-algorithm/bfs"
	"fmt"
)

func main() {
	grid := [][]bool{
		{false, false, false},
		{false, false, false},
		{false, false, false},
	}
	a := bfs.ShortestPath(grid, &bfs.Point{2, 0}, &bfs.Point{2, 2})
	fmt.Println(a)
}
