package main

import (
	binary_search "data-structure-and-algorithm/binary-search"
	"fmt"
)

func main() {
	res := binary_search.MountainSequence([]int{0, 1, 2, 1, 1, 1, 1, 1, 0})
	fmt.Println(res)
}
