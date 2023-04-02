package main

import "fmt"

// 二分查找先找到目标数字，然后从目标索引前后遍历得到出现次数
func search(nums []int, target int) int {
	index := binarySearch(nums, target)
	if index == -1 {
		return 0
	}
	count := 0
	for i := index; i < len(nums) && nums[i] == target; i++ {
		count++
	}
	for i := index - 1; i >= 0 && nums[i] == target; i-- {
		count++
	}
	return count
}

// 二分查找
func binarySearch(nums []int, target int) (index int) {
	length := len(nums)
	left, right := 0, length-1
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}

func main() {
	search([]int{1}, 1)
	fmt.Println(binarySearch([]int{1}, 1))
}
