package main

import (
	binary_search "data-structure-and-algorithm/binary-search"
	"fmt"
)

func main() {
	res := binary_search.KClosestNumbers([]int{1, 2, 4, 6, 8}, 3, 4)
	fmt.Println(res)
}
