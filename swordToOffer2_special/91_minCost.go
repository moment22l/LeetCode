package main

// dp[i][j]表示涂刷第i间房子为j颜色后的最小价值
// 再用滚动数组优化将空间降为O(1)
// 时间O(n)
//func minCost(costs [][]int) int {
//	var dp [2][3]int
//	for i := 0; i < 3; i++ {
//		dp[1][i] = costs[0][i]
//	}
//	for _, cost := range costs[1:] {
//		next := []int{min91(dp[1][1], dp[1][2]) + cost[0], min91(dp[1][0], dp[1][2]) + cost[1], min91(dp[1][0], dp[1][1]) + cost[2]}
//		for i := 0; i < 3; i++ {
//			dp[0][i], dp[1][i] = dp[1][i], next[i]
//		}
//	}
//	return min91(min91(dp[1][0], dp[1][1]), dp[1][2])
//}

func minCost(costs [][]int) int {
	dp := costs[0]
	for _, cost := range costs[1:] {
		dpNew := make([]int, 3)
		for j, c := range cost {
			dpNew[j] = min91(dp[(j+1)%3], dp[(j+2)%3]) + c
		}
		dp = dpNew
	}
	return min91(dp[0], min91(dp[1], dp[2]))
}

func min91(x, y int) int {
	if x < y {
		return x
	}
	return y
}
