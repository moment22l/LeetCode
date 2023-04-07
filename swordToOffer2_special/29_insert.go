package main

func insert(aNode *ListNode, x int) *ListNode {
	if aNode == nil {
		aNode = &ListNode{x, nil}
		aNode.Next = aNode
		return aNode
	}
	front, back := aNode, aNode
	front = front.Next
	for front != back {
		if front.Val >= x && x >= back.Val {
			back.Next = &ListNode{x, front}
			return aNode
		}
		if back.Val > front.Val && (front.Val >= x || back.Val <= x) {
			back.Next = &ListNode{x, front}
			return aNode
		}
		if front != aNode {
			front = front.Next
		}
		back = back.Next
	}
	aNode.Next = &ListNode{x, aNode.Next}
	return aNode
}
