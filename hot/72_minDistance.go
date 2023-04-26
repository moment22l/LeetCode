package main

// 时间O(mn) 空间O(mn)
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 0
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] != word2[j-1] {
				dp[i][j] = threeMin72(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
			} else {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

func threeMin72(i int, i2 int, i3 int) int {
	minNum := i
	if i2 < minNum {
		minNum = i2
	}
	if i3 < minNum {
		minNum = i3
	}
	return minNum
}

// dp[i][j] 表示 用word1前i个字符转化为word2前j个字符的最少操作次数
// 初始化:
// dp[0][0] = 0
// dp[0][j] = j, j > 0
// dp[i][0] = i, i > 0
// m=len(word1), n=len(word2)
// 状态转移方程:
// word1[i-1] != word2[j-1]时:
// dp[i][j] = min{dp[i-1][j-1], dp[i-1][j], dp[i][j-1]}+1, k∈[0, m]
// word1[i-1] == word2[j-1]时:
// dp[i][j] = dp[i-1][j-1]
