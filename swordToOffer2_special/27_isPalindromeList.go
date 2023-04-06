package main

//func isPalindromeList(head *ListNode) bool {
//	// 用数组记录链表的值, 再判断数组是否为回文数组
//	arr := make([]int, 0)
//	p := head
//	for p != nil {
//		arr = append(arr, p.Val)
//		p = p.Next
//	}
//	for i := 0; i < len(arr)/2; i++ {
//		if arr[i] != arr[len(arr)-i-1] {
//			return false
//		}
//	}
//	return true
//}

func isPalindromeList(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// 找到链表中点
	mid := findMiddleNode2(head)
	// 翻转右半边的链表
	tail := copyReverseList2(mid.Next)

	// 左半右半同时遍历比较
	front, back := head, tail
	result := true
	for back != nil {
		if front.Val != back.Val {
			result = false
			break
		}
		front = front.Next
		back = back.Next
	}
	// 将链表还原
	mid.Next = copyReverseList2(tail)
	return result
}

func findMiddleNode2(head *ListNode) *ListNode {
	// 找链表中点(快慢指针版)
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func copyReverseList2(head *ListNode) *ListNode {
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
