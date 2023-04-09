package main

import "LeetCode/utils"

func reorderList(head *utils.ListNode) {
	// 重排链表
	if head == nil {
		return
	}
	p1 := head
	p2 := findMiddleNode(head)
	next := p2.Next
	p2.Next = nil
	p2 = next
	p2 = copyReverseList(p2)
	if p1 == p2 || p2 == nil {
		return
	}
	mergeList(p1, p2)
}

func findMiddleNode(head *utils.ListNode) *utils.ListNode {
	// 找链表中点
	p := head
	length := 0
	for p != nil {
		length++
		p = p.Next
	}
	p = head
	for i := 0; i < length/2-1; i++ {
		p = p.Next
	}
	return p
}

func copyReverseList(head *utils.ListNode) *utils.ListNode {
	// 三指针遍历(前中后指针)
	var p1 *utils.ListNode
	p2 := head
	for p2 != nil {
		next := p2.Next
		p2.Next = p1
		p1 = p2
		p2 = next
	}
	return p1
}

func mergeList(p1 *utils.ListNode, p2 *utils.ListNode) {
	// 合并链表
	temp := p2
	for p1 != nil {
		next := p1.Next
		p1.Next = p2
		p1 = next
		next = p2.Next
		p2.Next = p1
		temp = p2
		p2 = next
	}
	if p2 != nil {
		temp.Next = p2
	}
}
