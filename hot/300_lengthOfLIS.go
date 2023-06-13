package main

// dp[i]表示以第i个元素为结尾的最长严格递增子序列的长度
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	ans := 0
	for i := 1; i < n; i++ {
		m := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j] > m {
				m = dp[j]
			}
		}
		dp[i] = m + 1
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
