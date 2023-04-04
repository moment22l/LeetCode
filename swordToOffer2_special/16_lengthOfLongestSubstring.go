package main

//func lengthOfLongestSubstring(s string) int {
//	// 哈希表存储每个字符最后出现的索引位置
//	m := make(map[byte]int)
//	pre := 0
//	max := 0
//	i := 0
//	for ; i < len(s); i++ {
//		if index, ok := m[s[i]]; ok {
//			if max < i-pre {
//				max = i - pre
//			}
//			if pre < index+1 {
//				pre = index + 1
//			}
//		}
//		m[s[i]] = i
//	}
//	if max < i-pre {
//		max = i - pre
//	}
//	return max
//}

func lengthOfLongestSubstring(s string) (ans int) {
	// 用数组(ascii码只有0-255)代替哈希表
	if len(s) == 0 {
		return 0
	}
	var m [256]int
	var countDup bool
	ans = 1
	left, right := 0, 0
	for ; right < len(s); right++ {
		m[s[right]]++
		if m[s[right]] == 2 {
			countDup = !countDup
		}
		for countDup {
			m[s[left]]--
			if m[s[left]] == 1 {
				countDup = !countDup
			}
			left++
		}
		if ans < right-left+1 {
			ans = right - left + 1
		}
	}
	return
}
