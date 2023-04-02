package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getLength(head *ListNode) int {
	length := 0
	for ; head != nil; head = head.Next {
		length++
	}
	return length
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{0, head}
	temp := dummy
	for i := 0; i < length-n; i++ {
		temp = temp.Next
	}
	temp.Next = temp.Next.Next
	return dummy.Next
}

func ShowNode(p *ListNode) { //遍历
	for p != nil {
		fmt.Println(*p)
		p = p.Next //移动指针
	}
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	head = removeNthFromEnd(head, 2)
	fmt.Println("1:")
	ShowNode(head)
	head = &ListNode{1, &ListNode{2, nil}}
	head = removeNthFromEnd(head, 1)
	fmt.Println("2:")
	ShowNode(head)
	head = &ListNode{1, nil}
	head = removeNthFromEnd(head, 1)
	fmt.Println("3:")
	ShowNode(head)
}
