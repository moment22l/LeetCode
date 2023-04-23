package main

import "LeetCode/utils"

func addTwoNumbers(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	head := new(utils.ListNode)
	node := new(utils.ListNode)
	carry := 0
	var val int
	sum := 0
	for {
		sum = 0
		if l1 != nil {
			sum = sum + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum = sum + l2.Val
			l2 = l2.Next
		}
		sum = sum + carry
		val = sum % 10
		carry = (sum - val) / 10
		if head.Next == nil {
			head.Val = val
		} else {
			node.Val = val
		}
		if l1 == nil && l2 == nil && carry == 0 {
			return head
		}
		if head.Next == nil {
			head.Next = node
		} else {
			node.Next = new(utils.ListNode)
			node = node.Next
		}
	}
}
