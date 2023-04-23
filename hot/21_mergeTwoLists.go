package main

import "LeetCode/utils"

func mergeTwoLists(list1 *utils.ListNode, list2 *utils.ListNode) *utils.ListNode {
	head := &utils.ListNode{
		Val:  0,
		Next: nil,
	}
	p := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = list1
			list1 = list1.Next
			p = p.Next
		} else {
			p.Next = list2
			list2 = list2.Next
			p = p.Next
		}
	}
	if list1 != nil {
		p.Next = list1
	}
	if list2 != nil {
		p.Next = list2
	}
	return head.Next
}
