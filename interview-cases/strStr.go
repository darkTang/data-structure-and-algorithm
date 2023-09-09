package interview_cases

import (
	"math"
)

/*
	暴力匹配算法
		1. 双重for循环
		2. 时间复杂度O(n^2)
*/

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	var hIndex, nIndex int
	for hIndex <= len(haystack)-len(needle) {
		for nIndex < len(needle) {
			if haystack[hIndex+nIndex] != needle[nIndex] {
				nIndex = 0
				break
			}
			nIndex++
			if nIndex == len(needle) {
				return hIndex
			}
		}
		hIndex++
	}
	return -1
}

/*
	Rabin Karp算法（https://coolcao.com/2020/08/20/rabin-karp/）
		1. 将字符串匹配转化为数值匹配，因此字符串匹配是逐个字符匹配，而hashcode是值匹配，降低时间复杂度
*/

func StrStr2(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	var BASE int = int(math.Pow10(6))
	var power int = 1
	var m = len(needle)
	for i := 0; i < m; i++ {
		power = (power * 26) % BASE
	}
	var needleCode = 0
	for i := 0; i < m; i++ {
		needleCode = (needleCode*26 + int(needle[i])) % BASE
	}

	var haystackCode = 0
	for i := 0; i < len(haystack); i++ {
		haystackCode = (haystackCode*26 + int(haystack[i])) % BASE
		if i < m-1 {
			continue
		}
		if i >= m {
			haystackCode = haystackCode - (int(haystack[i-m])*power)%BASE
			if haystackCode < 0 {
				haystackCode += BASE
			}
		}
		// double-check the string
		if haystackCode == needleCode {
			if haystack[i-m+1:i+1] == needle {
				return i - m + 1
			}
		}
	}
	return -1
}
