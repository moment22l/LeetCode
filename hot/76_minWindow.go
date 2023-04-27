package main

import "math"

// 双指针, 滑动窗口
// 时间O(k*m+n) 空间O(k) k为字符集长度, m为s的长度, n为t的长度
// 在本解答中为k=58, 实际只需要k=52(所有字母), 但为了方便操作直接申请了len=58的数组
func minWindow(s string, t string) (ans string) {
	m := len(s)
	var countT, countS [58]int
	for _, b := range t {
		countT[b-'A']++
	}
	minLength := math.MaxInt
	left, right := 0, 0
	for right < m {
		countS[s[right]-'A']++
		right++
		if isContain(countS, countT) {
			for left < right && isContain(countS, countT) {
				countS[s[left]-'A']--
				left++
			}
			left--
			countS[s[left]-'A']++
			if right-left < minLength {
				ans = s[left:right]
				minLength = right - left
			}
			if left < right {
				countS[s[left]-'A']--
				left++
			}
		}
	}
	return
}

func isContain(a, b [58]int) bool {
	for i := 0; i < 58; i++ {
		if a[i] < b[i] {
			return false
		}
	}
	return true
}
