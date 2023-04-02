package main

import "fmt"

func exchange(nums []int) []int {
	// 首尾双指针
	//left, right := 0, len(nums)-1
	//for left < right {
	//	for left < len(nums) && nums[left]%2 == 1 {
	//		left++
	//	}
	//	for right >= 0 && nums[right]%2 == 0 {
	//		right--
	//	}
	//	if left < right {
	//		nums[left], nums[right] = nums[right], nums[left]
	//	}
	//}
	//return nums

	// 快慢指针
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast]&1 == 1 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
	return nums
}

func main() {
	nums := []int{1, 8, 7, 1, 2, 3, 4, 5}
	fmt.Println(exchange(nums))
}
