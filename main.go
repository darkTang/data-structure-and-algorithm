package main

import (
	"data-structure-and-algorithm/bfs"
	"fmt"
)

func main() {
	a := bfs.AlienOrder([]string{"z", "x", "a", "zb", "zx"})
	fmt.Println(a)
}
