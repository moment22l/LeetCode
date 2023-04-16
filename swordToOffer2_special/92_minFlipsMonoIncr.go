package main

// dp[0]表示: 若将当前字符置0且使前面字符串满足要求, 所需要的翻转次数
// dp[1]同理
// 时间O(n) 空间O(1)
func minFlipsMonoIncr(s string) int {
	dp := make([]int, 2)
	if s[0] == '0' {
		dp[0], dp[1] = 0, 1
	} else if s[0] == '1' {
		dp[0], dp[1] = 1, 0
	}

	for _, b := range s[1:] {
		if b == '0' {
			dp[1] = min92(dp[0], dp[1]) + 1
		} else if b == '1' {
			dp[0], dp[1] = dp[0]+1, min92(dp[0], dp[1])
		}
	}
	return min92(dp[0], dp[1])
}

func min92(x, y int) int {
	if x < y {
		return x
	}
	return y
}
