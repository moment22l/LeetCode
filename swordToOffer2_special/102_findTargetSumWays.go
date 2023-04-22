package main

// 时间空间均为O(mn) m=len(nums), n=neg=(sum-target)/2
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum-target < 0 || (sum-target)%2 != 0 {
		return 0
	}
	neg := (sum - target) / 2
	dp := make([][]int, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		dp[i] = make([]int, neg+1)
	}
	dp[0][0] = 1
	for i, num := range nums {
		for j := 0; j <= neg; j++ {
			dp[i+1][j] = dp[i][j]
			if j >= num {
				dp[i+1][j] += dp[i][j-num]
			}
		}
	}
	return dp[len(nums)][neg]
}
