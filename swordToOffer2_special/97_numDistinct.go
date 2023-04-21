package main

// 当s[i-1]与t[j-1]相等时, dp[i][j] = dp[i-1][j-1]+dp[i-1][j], 前者为选上s[i-1]的情况, 后者为不选s[i-1]的情况
// 当s[i-1]与t[j-1]不相等时, dp[i][j] = dp[i-1][j]
// 时间O(mn) 空间O(mn) m=len1 n=len2
// 这里做了优化 空间到O(n) n=len2
// 相应的:
// 当s[i-1]与t[i-1]相等时, dp[j] += dp[j-1]
// 当s[i-1]与t[i-1]不相等时, dp[j]不变
// 注意 j 要从后往前遍历
func numDistinct(s string, t string) int {
	if len(t) == 0 {
		return 1
	}
	len1 := len(s)
	len2 := len(t)
	if len1 < len2 {
		return 0
	}
	dp := make([]int, len2+1)
	dp[0] = 1
	for i := 1; i <= len1; i++ {
		j := len2
		if i < len2 {
			j = i
		}
		for ; j > 0; j-- {
			if s[i-1] == t[j-1] {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[len2]
}
