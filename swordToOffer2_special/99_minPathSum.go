package main

// dp[j] = min{dp[j-1], dp[j]} + grid[i][j]
// 时间O(mn) 空间O(n)
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([]int, n)
	dp[0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		dp[0] += grid[i][0]
		for j := 1; j < n; j++ {
			dp[j] = min99(dp[j-1], dp[j]) + grid[i][j]
		}
	}
	return dp[n-1]
}

func min99(x, y int) int {
	if x < y {
		return x
	}
	return y
}
