package main

import "LeetCode/utils"

func insert(aNode *utils.ListNode, x int) *utils.ListNode {
	if aNode == nil {
		aNode = &utils.ListNode{x, nil}
		aNode.Next = aNode
		return aNode
	}
	front, back := aNode, aNode
	front = front.Next
	for front != back {
		if front.Val >= x && x >= back.Val {
			back.Next = &utils.ListNode{x, front}
			return aNode
		}
		if back.Val > front.Val && (front.Val >= x || back.Val <= x) {
			back.Next = &utils.ListNode{x, front}
			return aNode
		}
		if front != aNode {
			front = front.Next
		}
		back = back.Next
	}
	aNode.Next = &utils.ListNode{x, aNode.Next}
	return aNode
}
