package main

import (
	"data-structure-and-algorithm/sort"
	"fmt"
	"math/rand"
)

func main() {
	arr := make([]int, 100)
	for i := 0; i < 100; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(sort.MergeSort(arr))
}
