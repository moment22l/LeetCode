package main

// 时间O(mn) 空间O(mn) m=len1, n=len2
func isInterleave(s1 string, s2 string, s3 string) bool {
	len1 := len(s1)
	len2 := len(s2)
	if len1+len2 != len(s3) {
		return false
	}
	dp := make([][]bool, len1+1)
	for i := range dp {
		dp[i] = make([]bool, len2+1)
	}
	// 初始化
	dp[0][0] = true
	for i := 0; i < len1; i++ {
		dp[i+1][0] = dp[i][0] && s1[i] == s3[i]
		if !dp[i+1][0] {
			break
		}
	}
	for j := 0; j < len2; j++ {
		dp[0][j+1] = dp[0][j] && s2[j] == s3[j]
		if !dp[0][j+1] {
			break
		}
	}
	// 迭代
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			dp[i+1][j+1] = (dp[i][j+1] && s1[i] == s3[i+j+1]) || (dp[i+1][j] && s2[j] == s3[i+j+1])
		}
	}
	return dp[len1][len2]
}
