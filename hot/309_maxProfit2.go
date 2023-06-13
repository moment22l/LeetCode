package main

func maxProfit2(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	dp := make([][2]int, len(prices))
	dp[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		if i >= 2 {
			dp[i][0] = max309(dp[i-1][0], dp[i-2][1]-prices[i])
		} else {
			dp[i][0] = max309(-prices[0], -prices[1])
		}
		dp[i][1] = max309(dp[i-1][0]+prices[i], dp[i-1][1])
	}

	return dp[n-1][1]
}

func max309(x, y int) int {
	if x > y {
		return x
	}
	return y
}
