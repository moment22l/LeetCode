package main

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	// 暴力解法
	// var nums []int
	// for head != nil {
	// 	if len(nums) == 0 {
	// 		nums = append(nums, head.Val)
	// 	} else {
	// 		nums = append([]int{head.Val}, nums...)
	// 	}
	// 	head = head.Next
	// }
	// return nums

	// 栈
	//var nums1 []int
	//var nums2 []int
	//
	//for head != nil {
	//	nums1 = append(nums1, head.Val)
	//	head = head.Next
	//}
	//
	//for i := len(nums1) - 1; i >= 0; i-- {
	//	nums2 = append(nums2, nums1[i])
	//}
	//return nums2

	// 数组逆转1
	//var nums []int
	//
	//for head != nil {
	//	nums = append(nums, head.Val)
	//	head = head.Next
	//}
	//
	//n := len(nums)
	//for i := 0; i < n/2; i++ {
	//	nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	//}
	//return nums

	// 数组逆转2
	nums := make([]int, 0, 16)

	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
	return nums
}

func main() {
	head := &ListNode{0, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}}
	fmt.Println(reversePrint(head))
}
