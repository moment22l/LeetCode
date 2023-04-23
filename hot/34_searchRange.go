package main

import (
	"sort"
)

// 利用二分查找找到其中一个target，再向两边进行中心扩展
func searchRange(nums []int, target int) []int {
	if index, ok := binarySearch(nums, target); ok {
		// 当二分查找能找到target时，向左右延展找到边界值
		for left, right := index, index; ; {
			// 如果nums[left-1] == target就继续向左延展
			if left > 0 && nums[left-1] == target {
				left--
			}
			// 如果nums[right+1] == target 就继续向右延展
			if right < len(nums)-1 && nums[right+1] == target {
				right++
			}
			if (left == 0 || nums[left-1] != target) && (right == len(nums)-1 || nums[right+1] != target) {
				return []int{left, right}
			}
		}
	}
	return []int{-1, -1}
}

// 二分查找
func binarySearch(nums []int, target int) (int, bool) {
	var middle int
	for left, right := 0, len(nums)-1; left <= right; {
		middle = (left + right) / 2
		if nums[middle] == target {
			return middle, true
		}
		if nums[middle] <= target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1, false
}

// 利用API做二分查找
// sort.SearchInts寻找有序数组中第一个目标值出现的索引，无论目标值是否存在都会返回一个索引
// 如果目标值不存在，则会返回比目标值大的第一个位置的索引值
func searchRange2(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}
