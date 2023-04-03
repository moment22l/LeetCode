package main

import "sort"

func threeSum(nums []int) [][]int {
	// 先排序, 固定一个i, 然后用题6中双指针twoSum的变型找出三数之和为0的情况, 注意跳过数字相等的情况以去重三元组
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	output := make([][]int, 0)
	i := 0
	for i < len(nums)-2 {
		specialTwoSum(nums, i, &output)
		temp := nums[i]
		for nums[i] == temp && i < len(nums)-2 {
			i++
		}
	}
	return output
}

func specialTwoSum(nums []int, i int, result *[][]int) {
	j, k := i+1, len(nums)-1
	for j != k {
		if nums[i]+nums[j]+nums[k] == 0 {
			*result = append(*result, []int{nums[i], nums[j], nums[k]})
			temp := nums[j]
			for nums[j] == temp && j < k {
				j++
			}
		} else if nums[i]+nums[j]+nums[k] < 0 {
			j++
		} else {
			k--
		}
	}
}
