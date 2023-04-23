package main

// 1. 将两个数组按正序合并为一个nums, 然后直接找中位数
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	// 结果正确，但超时,且占用内存较多
//	var median float64
//	var nums []int
//	i := 0
//	j := 0
//	for {
//		if i != len(nums1) && j != len(nums2) {
//			if nums1[i] <= nums2[j] {
//				nums = append(nums, nums1[i])
//				i++
//			} else {
//				nums = append(nums, nums2[j])
//				j++
//			}
//		} else if i == len(nums1) && j != len(nums2) {
//			nums = append(nums, nums2[j])
//			j++
//		} else if i != len(nums1) && j == len(nums2) {
//			nums = append(nums, nums1[i])
//			i++
//		} else {
//			break
//		}
//	}
//	if len(nums)%2 == 0 {
//		median = (float64(nums[len(nums)/2-1]) + float64(nums[(len(nums)/2)])) / 2
//	} else {
//		median = float64(nums[len(nums)/2])
//	}
//	return median
//}

// 2. 比较 i+j 和 (len(nums1)+len(nums2))/2, 分情况判断哪个为中位数

// 3. 官方题解
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return 0
}
