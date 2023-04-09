package main

import "LeetCode/utils"

// 链表中环的入口节点

func detectCycle(head *utils.ListNode) *utils.ListNode {
	// 快慢指针
	slow, quick := head, head
	for quick != nil {
		slow = slow.Next
		if quick.Next == nil {
			return nil
		}
		quick = quick.Next.Next
		if slow == quick {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
