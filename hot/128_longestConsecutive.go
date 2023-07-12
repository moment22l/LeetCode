package main

import "sort"

//func longestConsecutive(nums []int) (ans int) {
//	n := len(nums)
//	numSet := make(map[int]bool)
//	for i := 0; i < n; i++ {
//		numSet[nums[i]] = true
//	}
//	for _, num := range nums {
//		if !numSet[num-1] {
//			currentNum := num
//			currentLength := 1
//			for numSet[currentNum+1] {
//				currentNum++
//				currentLength++
//			}
//			if currentLength > ans {
//				ans = currentLength
//			}
//		}
//	}
//	return ans
//}

func longestConsecutive(nums []int) (ans int) {
	n := len(nums)
	if n == 0 {
		return 0
	}
	sort.Ints(nums)
	length := 1
	for i := 1; i < n; i++ {
		if nums[i-1] == nums[i] {
			continue
		}
		if nums[i-1] == nums[i]-1 {
			length++
		} else {
			if ans < length {
				ans = length
			}
			length = 1
		}
	}
	if ans < length {
		ans = length
	}
	return ans
}
