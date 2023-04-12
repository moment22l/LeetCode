package linkedList

import "LeetCode/utils"

// SinglyGenerator 由数组生成单向链表
func SinglyGenerator(array []int) *utils.ListNode {
	dummy := &utils.ListNode{}
	temp := dummy
	for _, val := range array {
		temp.Next = &utils.ListNode{Val: val}
		temp = temp.Next
	}
	return dummy.Next
}

// DoublyGenerator 由数组生成双向链表
func DoublyGenerator(array []int) *utils.Node {
	dummy := &utils.Node{}
	temp := dummy
	for _, val := range array {
		newNode := &utils.Node{
			Val:  val,
			Prev: temp,
		}
		temp.Next = newNode
		temp = temp.Next
	}
	dummy.Next.Prev = nil
	return dummy.Next
}
