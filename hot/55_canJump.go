package main

// 动态规划
// dp[i] = nums[j]>=i-j && dp[j], j∈[i-maxJump，i-1], maxJump为前面的最大跳数
// 时间O(n^2) 空间O(n)
//func canJump(nums []int) bool {
//	n := len(nums)
//	dp := make([]bool, n)
//	dp[0] = true
//	maxJump := nums[0]
//	for i := 1; i < n; i++ {
//		for j := i - 1; j >= i-maxJump && j >= 0; j-- {
//			if nums[j] >= i-j && dp[j] {
//				dp[i] = true
//				break
//			}
//		}
//		if nums[i] > maxJump {
//			maxJump = nums[i]
//		}
//	}
//	return dp[n-1]
//}

// 动态规划 dp表示当前位置能跳到的最远距离
// dp[i] = max{dp[i-1], i+nums[i]}
// 时间O(n) 空间O(1)
func canJump(nums []int) bool {
	n := len(nums)
	dp := 0
	for i := 0; i < n; i++ {
		if i <= dp {
			dp = max55(dp, i+nums[i])
			if dp >= n-1 {
				return true
			}
			continue
		}
		break
	}
	return false
}

func max55(x, y int) int {
	if x > y {
		return x
	}
	return y
}
