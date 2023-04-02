package main

// 动态规划
// 初始化: 最右边和最下边的格子的最大价值固定
// 状态转移方程: dp[i][j] = max(dp[i+1][j], dp[i][j+1]) + grid[i][j]
func maxValue(grid [][]int) int {
	h, w := len(grid), len(grid[0])
	dp := make([][]int, h)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, w)
	}
	dp[h-1][w-1] = grid[h-1][w-1]
	for i := h - 2; i >= 0; i-- {
		dp[i][w-1] = dp[i+1][w-1] + grid[i][w-1]
	}
	for j := w - 2; j >= 0; j-- {
		dp[h-1][j] = dp[h-1][j+1] + grid[h-1][j]
	}
	for i := h - 2; i >= 0; i-- {
		for j := w - 2; j >= 0; j-- {
			dp[i][j] = max(dp[i+1][j], dp[i][j+1]) + grid[i][j]
		}
	}
	return dp[0][0]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
