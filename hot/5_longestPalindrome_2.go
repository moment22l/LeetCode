package main

// 中心扩展算法
func longestPalindrome(s string) string {
	if s == "" {
		return s
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		// 扩展中心为长度为1的子串
		left1, right1 := expandAroundCenter(s, i, i)
		// 扩展中心为长度为2的子串
		left2, right2 := expandAroundCenter(s, i, i+1)
		// 更新边界
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

// 从中心开始扩展，直至非回文子串出现
func expandAroundCenter(s string, left int, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
