package main

import (
	"LeetCode/utils"
	"fmt"
)

func getLength(head *utils.ListNode) int {
	length := 0
	for ; head != nil; head = head.Next {
		length++
	}
	return length
}

func removeNthFromEnd(head *utils.ListNode, n int) *utils.ListNode {
	length := getLength(head)
	dummy := &utils.ListNode{0, head}
	temp := dummy
	for i := 0; i < length-n; i++ {
		temp = temp.Next
	}
	temp.Next = temp.Next.Next
	return dummy.Next
}

func ShowNode(p *utils.ListNode) { //遍历
	for p != nil {
		fmt.Println(*p)
		p = p.Next //移动指针
	}
}
