package main

// 删除链表的倒数第 n 个结点

//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	// 先统计长度再找到倒数第n+1个, 然后删除第n个
//	// 特殊情况: 删除的是第一个时，判断如果length==n就直接返回head.Next
//	temp := head
//	length := 0
//	for temp != nil {
//		temp = temp.Next
//		length++
//	}
//	if length == n {
//		return head.Next
//	}
//	temp = head
//	for length > n+1 {
//		temp = temp.Next
//		length--
//	}
//	temp.Next = temp.Next.Next
//	return head
//}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 前后指针
	slow, quick := head, head
	for i := 0; i < n; i++ {
		quick = quick.Next
	}
	if quick == nil {
		return head.Next
	}
	for quick.Next != nil {
		slow = slow.Next
		quick = quick.Next
	}
	slow.Next = slow.Next.Next
	return head
}
