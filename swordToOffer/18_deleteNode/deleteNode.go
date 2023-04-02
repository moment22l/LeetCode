package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	// 双指针遍历
	//if head.Val == val {
	//	return head.Next
	//}
	//var prev, now *ListNode = head, head.Next
	//for now != nil && now.Val != val {
	//	now = now.Next
	//	prev = prev.Next
	//}
	//if now != nil {
	//	pre.Next = now.Next
	//}
	//return head

	// 单指针
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	cur := head
	for cur.Next != nil && cur.Next.Val != val {
		cur = cur.Next
	}
	if cur.Next != nil {
		cur.Next = cur.Next.Next
	}
	return head
}

func main() {

}
