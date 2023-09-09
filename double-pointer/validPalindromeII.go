package double_pointer

func validPalindrome(s string) bool {
	var left, right = 0, len(s) - 1
	for left < right {
		if s[left] != s[right] {
			break
		}
		left++
		right--
	}

	if left >= right {
		return true
	}

	return Palindrome(s, left+1, right) || Palindrome(s, left, right-1)
}

func Palindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 优化，公共逻辑提取为子函数
func validPalindrome2(s string) bool {
	var left, right = 0, len(s) - 1
	left, right = findDifference(s, left, right)

	if left >= right {
		return true
	}

	return Palindrome2(s, left+1, right) || Palindrome2(s, left, right-1)
}

func Palindrome2(s string, left, right int) bool {
	left, right = findDifference(s, left, right)
	return left >= right
}

func findDifference(s string, left, right int) (int, int) {
	for left < right && s[left] == s[right] {
		left++
		right--
	}
	return left, right
}
