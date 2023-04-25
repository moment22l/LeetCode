package main

//func maxSubArray(nums []int) int {
//	n := len(nums)
//	prefixSum := make([]int, n+1)
//	sum := 0
//	for i := 1; i <= n; i++ {
//		sum += nums[i-1]
//		prefixSum[i] = sum
//	}
//	minIndex := 0
//	max := math.MinInt
//	for i := 1; i <= n; i++ {
//		max = max53(max, prefixSum[i]-prefixSum[minIndex])
//		if prefixSum[i] < prefixSum[minIndex] {
//			minIndex = i
//		}
//	}
//	return max
//}

func maxSubArray(nums []int) int {
	n := len(nums)
	preSum := nums[0]
	max := nums[0]
	for i := 1; i < n; i++ {
		sum := max53(preSum+nums[i], nums[i])
		preSum = sum
		max = max53(max, sum)
	}
	return max
}

func max53(x, y int) int {
	if x > y {
		return x
	}
	return y
}
