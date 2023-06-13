package main

func numSquares(n int) int {
	squares := make([]int, 0)
	for i := 1; i*i <= n; i++ {
		squares = append(squares, i*i)
	}

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = n + 1
		for _, square := range squares {
			if i >= square {
				dp[i] = min279(dp[i], dp[i-square]+1)
			}
		}
	}
	return dp[n]
}

func min279(x, y int) int {
	if x < y {
		return x
	}
	return y
}
