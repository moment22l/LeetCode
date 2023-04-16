package main

// f(n) = max{f(n-2)+nums(n), f(n-1)}
// 时间O(n) 空间O(n)
//func rob(nums []int) int {
//	if len(nums) == 1 {
//		return nums[0]
//	}
//	n := len(nums)
//	dp := make([]int, n)
//	dp[0], dp[1] = nums[0], max89(nums[0], nums[1])
//	for i := 2; i < n; i++ {
//		dp[i] = max89(dp[i-2] + nums[i], dp[i-1])
//	}
//	return dp[n-1]
//}

// 滚动数组优化 空间降为O(1)
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp0, dp1 := nums[0], max89(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp0, dp1 = dp1, max89(dp0+nums[i], dp1)
	}
	return dp1
}

func max89(x, y int) int {
	if x > y {
		return x
	}
	return y
}
