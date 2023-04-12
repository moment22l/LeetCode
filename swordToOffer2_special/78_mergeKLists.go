package main

import "LeetCode/utils"

// 从头到尾两两合并 时间O(n*k^2) 空间O(1) n为lists的长度，n为单个链表的最长长度（可用分治来优化）
// 还有一种方法是优先队列，维护存储了所有链表当前遍历到的节点的优先队列，每次取优先队列的第一个
func mergeKLists(lists []*utils.ListNode) *utils.ListNode {
	if len(lists) == 0 {
		return nil
	}
	length := len(lists)
	for i := 1; i < length; i++ {
		newList := merge78(lists[0], lists[1])
		lists[1] = newList
		lists = lists[1:]
	}
	return lists[0]
}

func merge78(head1, head2 *utils.ListNode) *utils.ListNode {
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
