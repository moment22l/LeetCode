package main

// dp[i] = max{dp[i-1], dp[i-2]+nums[i]}
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	} else if n == 2 {
		return max198(nums[0], nums[1])
	}
	dp0, dp1 := nums[0], max198(nums[0], nums[1])
	for i := 2; i < n; i++ {
		maxDP := max198(dp1, dp0+nums[i])
		dp0, dp1 = dp1, maxDP
	}
	return dp1
}

func max198(x, y int) int {
	if x > y {
		return x
	}
	return y
}
