package main

// dp[i]表示到达第i个字符时所需要的分割次数
// 当0~i字符串为回文时,  dp[i] = 0
// 当0~i字符串不为回文时, dp[i] = min{dp[x-1]}+1, 其中x~i字符串为回文串
// 时间O(n^2) 空间O(n^2)
func minCut(s string) int {
	n := len(s)
	isPal := make([][]bool, n)
	for i := 0; i < n; i++ {
		isPal[i] = make([]bool, n)
		for j := 0; j <= i; j++ {
			if s[i] == s[j] && (i <= j+1 || isPal[j+1][i-1]) {
				isPal[j][i] = true
			}
		}
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if !isPal[i][0] {
			dp[i] = i
			for j := 1; j <= i; j++ {
				if isPal[i][j] {
					dp[i] = min94(dp[j-1]+1, dp[i])
				}
			}
		}

	}
	return dp[n-1]
}

func min94(x, y int) int {
	if x < y {
		return x
	}
	return y
}
