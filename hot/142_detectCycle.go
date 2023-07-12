package main

import "LeetCode/utils"

func detectCycle(head *utils.ListNode) *utils.ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for slow != p {
				slow = slow.Next
				p = p.Next
			}
			return p
		}
	}
	return nil
}
