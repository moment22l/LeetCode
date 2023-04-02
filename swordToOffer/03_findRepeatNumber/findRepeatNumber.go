package main

import (
	"fmt"
)

func findRepeatNumber(nums []int) int {
	// Map := make(map[int]int)
	// for _, num := range nums {
	// 	if _, ok := Map[num]; ok {
	// 		return num
	// 	}
	// 	Map[num] = num
	// }
	// return -1

	// n := len(nums)
	// temp := make([]int, n)
	// for i := range temp {
	// 	temp[i] = -1
	// }
	// for i := range nums {
	// 	if temp[nums[i]] == nums[i] {
	// 		return nums[i]
	// 	}
	// 	temp[nums[i]] = nums[i]
	// }
	// return -1

	i := 0
	for i < len(nums) {
		if nums[i] == i {
			i++
			continue
		}
		if nums[nums[i]] == nums[i] {
			return nums[i]
		}
		nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
	}
	return -1
}

func main() {
	fmt.Println(findRepeatNumber([]int{3, 4, 2, 0, 1, 0}))
}
