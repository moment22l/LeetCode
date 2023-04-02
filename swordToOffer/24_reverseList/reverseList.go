package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// 用栈反转
	//if head == nil {
	//	return nil
	//}
	//var stack []*ListNode
	//for head != nil {
	//	stack = append(stack, head)
	//	head = head.Next
	//}
	//stack, head = pop(stack)
	//var now, prev *ListNode
	//prev = head
	//for len(stack) > 0 {
	//	stack, now = pop(stack)
	//	prev.Next = now
	//	prev = prev.Next
	//}
	//if now != nil {
	//	now.Next = nil
	//}
	//return head

	// 双指针迭代(更优雅的写法)
	//var former *ListNode
	//latter := head
	//for latter != nil {
	//	temp := latter.Next
	//	latter.Next = former
	//	former = latter
	//	latter = temp
	//}
	//return former

	// 递归
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func pop(stack []*ListNode) ([]*ListNode, *ListNode) {
	s := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return stack, s
}

func main() {

}
