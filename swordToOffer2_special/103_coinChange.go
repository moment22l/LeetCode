package main

// dp[i]表示总额为i的最小硬币数目
// dp[i] = min{dp[i], dp[i-coin]+1}, coin range coins
// 时间O(nt) 空间O(t), n=len(coins), t=amount
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min103(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func min103(x, y int) int {
	if x < y {
		return x
	}
	return y
}
