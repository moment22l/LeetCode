package main

import "fmt"

func maxSubArray(nums []int) int {
	sum := nums[0]
	max := sum
	for i := 1; i < len(nums); i++ {
		if sum+nums[i] > nums[i] {
			sum = sum + nums[i]
		} else {
			sum = nums[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func main() {
	fmt.Println(maxSubArray([]int{-1, -1, -2, -2}))
}
