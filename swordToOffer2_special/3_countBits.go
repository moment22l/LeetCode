package main

// 	题目: 前 n 个数字二进制中 1 的个数
//	给定一个非负整数 n ，请计算 0 到 n 之间的每个数字的二进制表示中 1 的个数，并输出一个数组。

func countBits(n int) []int {
	nums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		nums[i] = nums[i&(i-1)] + 1
		// nums[i] = nums[i>>1] + (i&1)
	}
	return nums
}
