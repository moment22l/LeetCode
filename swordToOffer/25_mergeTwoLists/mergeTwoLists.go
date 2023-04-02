package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 直接遍历
	if l1 == nil && l2 == nil {
		return nil
	}
	head := &ListNode{0, nil}
	now := head
	for l1 != nil && l2 != nil {
		if l1 != nil && l2 == nil || l1.Val <= l2.Val {
			now.Next = l1
			l1 = l1.Next
			now = now.Next
		} else if l1 == nil && l2 != nil || l1.Val > l2.Val {
			now.Next = l2
			l2 = l2.Next
			now = now.Next
		}
	}
	if l1 != nil {
		now.Next = l1
	} else if l2 != nil {
		now.Next = l2
	}
	head = head.Next
	return head
}

func main() {

}
