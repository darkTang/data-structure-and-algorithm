package bfs

import "container/list"

func LadderLength(beginWord string, endWord string, wordList []string) int {
	if !isExist(endWord, wordList) {
		return 0
	}
	queue := list.New()
	visited := make(map[string]bool)
	distance := 0
	queue.PushBack(beginWord)
	visited[beginWord] = true
	var curWord string
	for queue.Len() > 0 {
		length := queue.Len()
		distance++
		for i := 0; i < length; i++ {
			curWord = queue.Remove(queue.Front()).(string)
			if curWord == endWord {
				return distance
			}
			for _, nextWord := range getNextWords(curWord, wordList) {
				if !visited[nextWord] {
					queue.PushBack(nextWord)
					visited[nextWord] = true
				}
			}
		}
	}
	return 0
}

func isExist(endWord string, wordList []string) bool {
	for _, word := range wordList {
		if word == endWord {
			return true
		}
	}
	return false
}

func getNextWords(curWord string, wordList []string) []string {
	res := make([]string, 0)
	for _, word := range wordList {
		if len(word) != len(curWord) {
			continue
		}
		length := 0
		for i := 0; i < len(curWord); i++ {
			if curWord[i] == word[i] {
				length++
			}
		}
		if length == len(curWord)-1 {
			res = append(res, word)
		}
	}
	return res
}
