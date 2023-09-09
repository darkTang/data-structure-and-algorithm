package double_pointer

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	var left, right = 0, len(s) - 1
	for left < right {
		for left < right && !isValid(rune(s[left])) {
			left++
		}
		for left < right && !isValid(rune(s[right])) {
			right--
		}
		if left < right && strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func isValid(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}
