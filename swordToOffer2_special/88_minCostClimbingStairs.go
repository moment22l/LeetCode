package main

// f(n) = min{f(n-1)+cost(n-1), f(n-2)+cost(n-2)}
// 时间O(n) 空间O(n)
//func minCostClimbingStairs(cost []int) int {
//	n := len(cost)
//	dp := make([]int, n+1)
//	dp[0], dp[1] = 0, 0
//	for i := 2; i <= n; i++ {
//		dp[i] = min88(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
//	}
//	return dp[n]
//}

// 优化: 滚动数组 空间降到O(1)
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp0, dp1 := 0, 0
	for i := 2; i <= n; i++ {
		dp0, dp1 = dp1, min88(dp1+cost[i-1], dp0+cost[i-2])
	}
	return dp1
}

func min88(x, y int) int {
	if x < y {
		return x
	}
	return y
}
