package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	node := new(ListNode)
	carry := 0
	var val int
	sum := 0
	for {
		sum = 0
		if l1 != nil {
			sum = sum + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum = sum + l2.Val
			l2 = l2.Next
		}
		sum = sum + carry
		val = sum % 10
		carry = (sum - val) / 10
		if head.Next == nil {
			head.Val = val
		} else {
			node.Val = val
		}
		if l1 == nil && l2 == nil && carry == 0 {
			return head
		}
		if head.Next == nil {
			head.Next = node
		} else {
			node.Next = new(ListNode)
			node = node.Next
		}
	}
}

func main() {
	l1 := new(ListNode)
	l2 := new(ListNode)
	l1.Val = 2
	l1.Next = new(ListNode)
	l1.Next.Val = 4
	l1.Next.Next = new(ListNode)
	l1.Next.Next.Val = 3
	l2.Val = 5
	l2.Next = new(ListNode)
	l2.Next.Val = 6
	l2.Next.Next = new(ListNode)
	l2.Next.Next.Val = 4
	total := addTwoNumbers(l1, l2)
	for total != nil {
		fmt.Println(total.Val)
		total = total.Next
	}
}
