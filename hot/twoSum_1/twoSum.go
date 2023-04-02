package main

import "fmt"

//func twoSum(nums []int, target int) []int {
//  // 最简单的双循环，时间复杂度为O(n^2)
//	var answer []int
//	for i := 0; i < len(nums); i++ {
//		for j := 0; j < len(nums); j++ {
//			if i == j {
//				continue
//			}
//			if nums[i]+nums[j] == target {
//				answer = append(answer, i)
//				answer = append(answer, j)
//				return answer
//			}
//		}
//	}
//	return nil
//}

func twoSum(nums []int, target int) []int {
	// 通过map记录nums中的数字及其索引，遍历寻找target-num[i]是否存在于map，存在即为答案
	// 时间复杂度为O(n)
	var answer []int
	numsMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := numsMap[target-nums[i]]; ok {
			answer = append(answer, numsMap[target-nums[i]])
			answer = append(answer, i)
			return answer
		}
		numsMap[nums[i]] = i
	}
	return nil
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	answer := twoSum(nums, target)
	fmt.Println(answer)
}
