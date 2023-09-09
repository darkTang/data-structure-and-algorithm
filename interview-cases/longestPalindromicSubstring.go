package interview_cases

/*
	暴力枚举算法
		1. 外层两层for循环用来找所有的连续子串
		2. 里层for循环用来判断是否是回文子串
		3. 时间复杂度为O(n^3)
		4. 优化：因为求的是最长回文子串，因此可以改变循环策略，按长度从大到小来遍历
*/

func LongestPalindrome1(s string) string {
	for length := len(s); length > 0; length-- {
		for start := 0; start+length <= len(s); start++ {
			if isPalindrome1(s[start : start+length]) {
				return s[start : start+length]
			}
		}
	}
	return ""
}

func isPalindrome1(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

/*
	中心线枚举算法
		x|a|b|c|b|a|c
		1. 从中间向两边枚举
		2. 双指针记录回文子串，找出最长回文子串
		3. 初始化双指针（字符串长度分为奇数或偶数）
		4. 时间复杂度O(n^2)
*/

func LongestPalindrome2(s string) string {
	var longest string
	for i := 0; i < len(s); i++ {
		oddPalindrome := getPalindromeFrom(s, i, i)
		if len(longest) < len(oddPalindrome) {
			longest = oddPalindrome
		}
		evenPalindrome := getPalindromeFrom(s, i, i+1)
		if len(longest) < len(evenPalindrome) {
			longest = evenPalindrome
		}
	}
	return longest
}

func getPalindromeFrom(s string, left, right int) string {
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			break
		}
		left--
		right++
	}
	return s[left+1 : right]
}

/*
	动态规划算法【区间型动态规划】
		1. 这里的动态规划会在后面上面玩后再来
*/

func LongestPalindrome3(s string) string {
	return ""
}
