package main

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	dp[0] = grid[0][0]
	for j := 1; j < n; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[j] += grid[i][j]
			} else {
				dp[j] = min64(dp[j-1], dp[j]) + grid[i][j]
			}
		}
	}
	return dp[n-1]
}

func min64(x, y int) int {
	if x < y {
		return x
	}
	return y
}
