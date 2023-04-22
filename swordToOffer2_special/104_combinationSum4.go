package main

// dp[i]表示总和为i的排列数目
// dp[i] = ∑dp[i-x], x为nums中所有元素构成的集合
// 时间O(nt) 空间O(t), n=len(nums), t=target
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i < target+1; i++ {
		for _, num := range nums {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

func min104(x, y int) int {
	if x < y {
		return x
	}
	return y
}
