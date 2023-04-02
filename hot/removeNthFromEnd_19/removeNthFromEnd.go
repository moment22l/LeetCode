package main

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// 利用递归回溯的方式找出需要删除的节点，类似官方题解中的方法二：利用栈来寻找目标节点的前驱节点
// 但不知为何在leetcode上提交时执行会错误（在本地执行没问题）
//var listLen int = -1 // 记录链表长度
//
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	listLen = -1
//	// 如果只有一个节点就直接返回空链表
//	if head.Next == nil {
//		head = nil
//		return head
//	}
//	dummy := &ListNode{0, head} // 哑结点
//	remove(dummy, n, -1)        // 删除对应节点
//	return dummy.Next
//}
//
//func remove(now *ListNode, n int, index int) {
//	listLen++ // 计算链表长度
//
//	// 遍历链表并使index++记录节点的编号
//	if now.Next != nil {
//		remove(now.Next, n, index+1)
//	}
//	// 递归过程中，找出对应节点的前驱节点，并使其Next指向对应节点的后继节点
//	if index == listLen-n-1 {
//		now.Next = now.Next.Next
//	}
//}
//
//func ShowNode(p *ListNode) { //遍历
//	for p != nil {
//		fmt.Println(*p)
//		p = p.Next //移动指针
//	}
//}
//
//func main() {
//	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
//	head = removeNthFromEnd(head, 2)
//	fmt.Println("1:")
//	ShowNode(head)
//	head = &ListNode{1, &ListNode{2, nil}}
//	head = removeNthFromEnd(head, 1)
//	fmt.Println("2:")
//	ShowNode(head)
//	head = &ListNode{1, nil}
//	head = removeNthFromEnd(head, 1)
//	fmt.Println("3:")
//	ShowNode(head)
//}
