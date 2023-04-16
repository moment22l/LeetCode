package main

// 用题89的方法分段求0-n-1和1-n两种情况即可
// f(n) = max{f(n-2)+nums(n), f(n-1)}
// 时间O(n) 空间O(1)
func circleRob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max90(nums[0], nums[1])
	}
	var dp [2][2]int
	dp[0][0], dp[0][1] = nums[0], max90(nums[0], nums[1])
	dp[1][0], dp[1][1] = nums[1], max90(nums[1], nums[2])
	for i := 3; i < len(nums); i++ {
		dp[0][0], dp[0][1] = dp[0][1], max90(dp[0][0]+nums[i-1], dp[0][1])
		dp[1][0], dp[1][1] = dp[1][1], max90(dp[1][0]+nums[i], dp[1][1])
	}
	return max90(dp[0][1], dp[1][1])
}

func max90(x, y int) int {
	if x > y {
		return x
	}
	return y
}
