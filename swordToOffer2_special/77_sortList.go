package main

import "LeetCode/utils"

// 快慢指针找中点，归并排序 时间O(nlogn) 空间O(logn)
func sortList(head *utils.ListNode) *utils.ListNode {
	return sort77(head, nil)
}

func mergeSortList(head1, head2 *utils.ListNode) *utils.ListNode {
	dummy := &utils.ListNode{}
	temp, temp1, temp2 := dummy, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummy.Next
}

func sort77(head, tail *utils.ListNode) *utils.ListNode {
	if head == nil {
		return nil
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	return mergeSortList(sort77(head, slow), sort77(slow, tail))
}
