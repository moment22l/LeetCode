package main

// 时间空间均为O(mn) m=len(nums), n=sum/2
// dp[i][j] = dp[i-1][j], dp[i-1][j]==true || j < nums[i-1]
// dp[i][j] = dp[i-1][j-nums[i-1]], dp[i-1][j]==false && j >= nums[i-1]
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	count := len(nums)
	dp := make([][]bool, count+1)
	for i := 0; i <= count; i++ {
		dp[i] = make([]bool, sum/2+1)
		dp[i][0] = true
	}
	for i := 1; i <= count; i++ {
		for j := 1; j <= sum/2; j++ {
			dp[i][j] = dp[i-1][j]
			if !dp[i][j] && j >= nums[i-1] {
				dp[i][j] = dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[count][sum/2]
}
