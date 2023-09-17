package main

import "fmt"

func minOperations(start, end int) int {
	if end < start {
		return 0
	}

	count := 0
	for end > start {
		if end%2 == 0 {
			end /= 2
		} else {
			end++
		}
		count++
	}

	return count
}

func main() {
	var start, end int
	fmt.Scanln(&start)
	fmt.Scanln(&end)
	fmt.Println(minOperations(start, end))
}
