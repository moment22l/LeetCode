package main

import "fmt"

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

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println("[1,2,3]: ", findMedianSortedArrays(nums1, nums2))
	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	fmt.Println("[1,2,3,4]: ", findMedianSortedArrays(nums1, nums2))
	nums1 = []int{}
	nums2 = []int{1}
	fmt.Println("[1]: ", findMedianSortedArrays(nums1, nums2))
	nums1 = []int{1, 3, 999, 7777, 9999}
	nums2 = []int{0, 2, 888, 6666, 8888}
	fmt.Println("[0,1,2,3,4,5,6,7,8,9]: ", findMedianSortedArrays(nums1, nums2))
}
