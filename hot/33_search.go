package main

// 二分查找
func search(nums []int, target int) int {
	var middle int
	for left, right := 0, len(nums)-1; left <= right; {
		middle = (left + right) / 2
		// 每次判断middle所指向的位置是否与target相等
		if nums[middle] == target {
			return middle
		}
		if nums[left] <= nums[middle] { // 1. nums[left:middle-1]为有序切片
			if nums[left] <= target && nums[middle] > target {
				// (1) target处于nums[left]~nums[middle-1]范围中，就到left~(middle-1)中继续寻找
				right = middle - 1
			} else {
				// (2) target 不处于nums[left]~nums[middle-1]范围中，就到(middle+1)~right中继续寻找
				left = middle + 1
			}
		} else { // 2. nums[left:middle]非有序切片，则nums[middle:right]为有序切片
			if nums[middle] < target && nums[right] >= target {
				// (1) target处于nums[middle+1]~nums[right-1]范围中，就到(middle+1)~right中继续寻找
				left = middle + 1
			} else {
				// (2) target不处于nums[middle+1]~nums[right-1]范围中，就到left~(middle-1)中继续寻找
				right = middle - 1
			}
		}
	}
	return -1
}
