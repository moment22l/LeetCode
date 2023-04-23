package main

import (
	"sort"
)

// 最暴力解法时间复杂度达到O(n^3)，超时
// 定住一个，然后左(i+1)右(len(nums)-1)双指针，向中间靠拢(O(n^2))
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	// 遍历数组
	for i := 0; i < len(nums); i++ {
		// 遍历过程中重复元素就跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 左右双指针，left指向i后第一个，right指向数组最后一个元素
		left, right := i+1, len(nums)-1
		// 将左指针每次向后移动一位
		for ; left < len(nums); left++ {
			// 重复元素跳过
			if left > i+1 && nums[left] == nums[left-1] {
				continue
			}
			// 如果左指针在右指针之前且三数之和大于0，则右指针向左移动一位
			for left < right && nums[i]+nums[left]+nums[right] > 0 {
				right--
			}
			// 当左指针与右指针相遇时，结束本次循环
			if left == right {
				break
			}
			// 判断三数之和是否为0，是则加入到答案数组中
			if nums[i]+nums[left]+nums[right] == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
			}
		}
	}
	return ans
}
