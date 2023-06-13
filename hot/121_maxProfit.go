package main

func maxProfit(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	dp := 0
	minP := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minP {
			minP = prices[i]
		} else if prices[i]-minP > dp {
			dp = prices[i] - minP
		}
	}
	return dp
}
