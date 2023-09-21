package bfs

import (
	"container/list"
)

func AlienOrder(words []string) string {
	graph := make(map[byte][]byte)
	inDegree := make(map[byte]int)
	for _, c := range words[0] {
		inDegree[byte(c)] = 0
	}
next:
	for i := 1; i < len(words); i++ {
		s, t := words[i-1], words[i]
		for _, c := range t {
			inDegree[byte(c)] += 0
		}
		for j := 0; j < len(s) && j < len(t); j++ {
			if s[j] != t[j] {
				graph[s[j]] = append(graph[s[j]], t[j])
				inDegree[t[j]]++
				continue next
			}
		}
		if len(s) > len(t) {
			return ""
		}
	}
	order := make([]byte, 0)
	queue := list.New()
	var curByte byte
	for key, value := range inDegree {
		if value == 0 {
			queue.PushBack(key)
		}
	}
	for queue.Len() > 0 {
		curByte = queue.Remove(queue.Front()).(byte)
		order = append(order, curByte)
		for _, v := range graph[curByte] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue.PushBack(v)
			}
		}
	}
	if len(order) != len(inDegree) {
		return ""
	}
	return string(order)
}
