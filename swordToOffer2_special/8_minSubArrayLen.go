package main

import "sort"

// func minSubArrayLen(target int, nums []int) int {
//  // 滑动窗口, 小了就让右边界往右扩, 大了就让左边界往右缩
//  i, j := 0, 0
//	min := 0
//	nowSum := nums[0]
//	for i <= j && j < len(nums) {
//		if nowSum >= target {
//			if min == 0 || min > j-i+1 {
//				min = j - i + 1
//			}
//			nowSum -= nums[i]
//			i++
//			continue
//		}
//		j++
//		if j < len(nums) {
//			nowSum += nums[j]
//		}
//	}
//	return min
//}

//func minSubArrayLen(target int, nums []int) int {
//	// 优化: 让left一次性右移到sum<target的位置
//	left := 0
//	min := 0
//	sum := 0
//	for right := 0; right < len(nums); right++ {
//		sum += nums[right]
//		for left <= right && sum >= target {
//			if min > right-left+1 || min == 0 {
//				min = right - left + 1
//			}
//			sum -= nums[left]
//			left++
//		}
//	}
//	return min
//}

func minSubArrayLen(target int, nums []int) int {
	// 官方题解: 前缀和+二分查找
	length := len(nums)
	if length == 0 {
		return 0
	}
	min := 0
	sum := make([]int, length+1)
	for i := 1; i < length+1; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}

	for i := 1; i < length+1; i++ {
		k := target + sum[i-1]
		bound := sort.SearchInts(sum, k)
		if bound <= length {
			if min > bound-i+1 || min == 0 {
				min = bound - i + 1
			}
		}
	}
	return min
}
