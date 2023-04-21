package main

// 时间O(mn) 空间O(mn) m=len1, n=len2
//func longestCommonSubsequence(text1 string, text2 string) int {
//	len1 := len(text1)
//	len2 := len(text2)
//	dp := make([][]int, len1+1)
//	for i := range dp {
//		dp[i] = make([]int, len2+1)
//	}
//	for i := 0; i < len1; i++ {
//		for j := 0; j < len2; j++ {
//			if text1[i] == text2[j] {
//				dp[i+1][j+1] = dp[i][j] + 1
//			} else {
//				dp[i+1][j+1] = max95(dp[i][j+1], dp[i+1][j])
//			}
//		}
//	}
//	return dp[len1][len2]
//}

// 优化空间至O(n), n=len2 ps: 还可以进一步优化成一个一维数组
func longestCommonSubsequence(text1 string, text2 string) int {
	len1 := len(text1)
	len2 := len(text2)
	dp := make([][]int, 2)
	dp[0], dp[1] = make([]int, len2+1), make([]int, len2+1)
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if text1[i] == text2[j] {
				dp[1][j+1] = dp[0][j] + 1
			} else {
				dp[1][j+1] = max95(dp[0][j+1], dp[1][j])
			}
		}
		copy(dp[0], dp[1])
	}
	return dp[0][len2]
}

func max95(x, y int) int {
	if x > y {
		return x
	}
	return y
}
