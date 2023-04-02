package main

import "fmt"

// 超时
//func longestPalindrome(s string) string {
//	maxString := ""
//	runes := []rune(s)
//	lastOccured := make(map[rune][]int)
//	flag := false
//	for i, ch := range runes {
//		if lastI, ok := lastOccured[ch]; ok {
//			for j := range lastI {
//				nowString := string(runes[lastI[j] : i+1])
//				if judgePalindrome(nowString) {
//					flag = true
//					if len(nowString) > len(maxString) {
//						maxString = nowString
//					}
//				}
//			}
//		}
//		lastOccured[ch] = append(lastOccured[ch], i)
//	}
//	if flag == false {
//		if len(runes) == 0 {
//			return ""
//		} else {
//			return string(runes[0])
//		}
//	}
//	return maxString
//}
//
//func judgePalindrome(s string) bool {
//	i, j := 0, len(s)-1
//	runes := []rune(s)
//	for {
//		if runes[i] != runes[j] {
//			return false
//		}
//		if i >= j {
//			return true
//		}
//		i++
//		j--
//	}
//}

// 动态规划算法
func longestPalindrome1(s string) string {
	len := len(s)
	if len < 2 {
		return s
	}

	// dp[i][j]表示s(i...j)
	maxLen := 1
	begin := 0
	// 初始化dp数组
	dp := make([][]bool, len)
	for i := range dp {
		dp[i] = make([]bool, len)
	}

	runes := []rune(s)
	// 将len为1的子串默认设为true
	for i := 0; i < len; i++ {
		dp[i][i] = true
	}

	for L := 2; L <= len; L++ {
		for i := 0; i < len; i++ {
			j := i + L - 1
			// 右边界越界则结束循环
			if j >= len {
				break
			}

			// 前后字母不等说明该子串不为回文子串
			if runes[i] != runes[j] {
				dp[i][j] = false
			} else {
				// 当子串长度小于等于3时，由于前后字母相同，所以一定为回文子串
				if j-i+1 <= 3 {
					dp[i][j] = true
				} else {
					// 前后字母一致，则由去掉前后字母的子串决定该子串是否为子串
					dp[i][j] = dp[i+1][j-1]
				}
			}

			// 更新回文子串的开头及长度
			if dp[i][j] && L > maxLen {
				maxLen = L
				begin = i
			}
		}
	}
	return string(runes[begin : begin+maxLen])
}

func main() {
	fmt.Println(longestPalindrome1("babad"))
	fmt.Println(longestPalindrome1("cbbd"))
	fmt.Println(longestPalindrome1(""))
	fmt.Println(longestPalindrome1("abcdef"))
	fmt.Println(longestPalindrome1("bbbbbb"))
	fmt.Println(longestPalindrome1("cbbcbbcbbcbbc"))
	fmt.Println(longestPalindrome1("abababababab"))
}
