package bfs

import (
	"container/list"
)

type Point struct {
	X, Y int
}

// ShortestPath 方法1 visited记录最短路径值
func ShortestPath(grid [][]bool, source *Point, destination *Point) int {
	if grid == nil || len(grid) == 0 || grid[0] == nil || len(grid[0]) == 0 {
		return -1
	}
	directs := [][]int{{1, 2}, {1, -2}, {-1, 2}, {-1, -2}, {2, 1}, {2, -1}, {-2, 1}, {-2, -1}}
	queue := list.New()
	visited := make(map[[2]int]int)
	queue.PushBack(source)
	// visited记录的是当前最短路径
	// 这样写就不要再加一层for循环来进行层序遍历
	visited[[2]int{source.X, source.Y}] = 0
	var curCoordinate *Point
	for queue.Len() > 0 {
		curCoordinate = queue.Remove(queue.Front()).(*Point)
		if curCoordinate.X == destination.X && curCoordinate.Y == destination.Y {
			return visited[[2]int{curCoordinate.X, curCoordinate.Y}]
		}
		for _, nextPoint := range getNextPoints(grid, curCoordinate, visited, directs) {
			queue.PushBack(nextPoint)
			visited[[2]int{nextPoint.X, nextPoint.Y}] = visited[[2]int{curCoordinate.X, curCoordinate.Y}] + 1
		}
	}
	return -1
}

func getNextPoints(grid [][]bool, curCoordinate *Point, visited map[[2]int]int, directs [][]int) (nextPoints []*Point) {
	for _, direct := range directs {
		var point *Point
		point = &Point{
			X: curCoordinate.X + direct[0],
			Y: curCoordinate.Y + direct[1],
		}
		if _, ok := visited[[2]int{point.X, point.Y}]; !ok && isValidPoint(grid, point) {
			nextPoints = append(nextPoints, point)
		}
	}
	return
}

func isValidPoint(grid [][]bool, point *Point) bool {
	lenX := len(grid)
	lenY := len(grid[0])
	if point.X < 0 || point.X >= lenX || point.Y < 0 || point.Y >= lenY || grid[point.X][point.Y] {
		return false
	}
	return true
}

// ShortestPath2 方法2  如果visited只记录是否访问过，那么则需要层序遍历
func ShortestPath2(grid [][]bool, source *Point, destination *Point) int {
	if grid == nil || len(grid) == 0 || grid[0] == nil || len(grid[0]) == 0 {
		return -1
	}
	directs := [][]int{{1, 2}, {1, -2}, {-1, 2}, {-1, -2}, {2, 1}, {2, -1}, {-2, 1}, {-2, -1}}
	distance := 0
	queue := list.New()
	visited := make(map[[2]int]bool)
	queue.PushBack(source)
	visited[[2]int{source.X, source.Y}] = true
	var curCoordinate *Point
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			curCoordinate = queue.Remove(queue.Front()).(*Point)
			//fmt.Println(distance)
			if curCoordinate.X == destination.X && curCoordinate.Y == destination.Y {
				return distance
			}
			for _, nextPoint := range getNextPoints2(grid, curCoordinate, visited, directs) {
				queue.PushBack(nextPoint)
				visited[[2]int{nextPoint.X, nextPoint.Y}] = true
			}
		}
		distance++
	}
	return -1
}

func getNextPoints2(grid [][]bool, curCoordinate *Point, visited map[[2]int]bool, directs [][]int) (nextPoints []*Point) {
	for _, direct := range directs {
		var point *Point
		point = &Point{
			X: curCoordinate.X + direct[0],
			Y: curCoordinate.Y + direct[1],
		}
		if !visited[[2]int{point.X, point.Y}] && isValidPoint2(grid, point) {
			nextPoints = append(nextPoints, point)
		}
	}
	return
}

func isValidPoint2(grid [][]bool, point *Point) bool {
	lenX := len(grid)
	lenY := len(grid[0])
	if point.X < 0 || point.X >= lenX || point.Y < 0 || point.Y >= lenY || grid[point.X][point.Y] {
		return false
	}
	return true
}

/*
在 Go 中，当一个指针类型作为 map 的 key 时，会导致内存消耗过大的主要原因是指针的比较操作。
在 map 中，为了判断两个 key 是否相等，会使用 == 运算符进行比较。对于指针类型的 key，比较的是指针的值，也就是内存地址。而指针的地址空间非常大，因此比较两个指针的值会导致更多的内存开销。
另外，由于指针可以指向任意类型的数据，为了确保指针作为 key 的正确性，需要比较指针所指向的实际值是否相等。这就要求在比较两个指针是否相等时，还需要对指针所指向的数据进行深层次的比较，这会进一步增加内存消耗。
因此，如果需要使用指针类型作为 map 的 key，并且关心内存消耗的问题，可以考虑使用其他可比较的类型作为 key，或者使用某种方式将指针类型转换为可比较的类型，例如将指针转换为字符串作为 key。
*/
