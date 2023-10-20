package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max_213(nums[0], nums[1])
	}
	return max_213(steal(nums[:len(nums)-1]), steal(nums[1:]))
}

func steal(nums []int) int {
	// f(i) = max{f(i-1), f(i-2)+nums[i]}
	a, b := nums[0], max_213(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		a, b = b, max_213(b, a+nums[i])
	}
	return b
}

func max_213(x, y int) int {
	if x > y {
		return x
	}
	return y
}
