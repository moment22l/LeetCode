package main

func reverseList(head *ListNode) *ListNode {
	// 三指针遍历(前中后指针)
	var p1 *ListNode
	p2 := head
	for p2 != nil {
		next := p2.Next
		p2.Next = p1
		p1 = p2
		p2 = next
	}
	return p1
}
