package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	// 两次遍历，第一次确定链表长度，第二次定位到倒数第k个节点并返回
	//n := 0
	//now := head
	//for now != nil {
	//	n++
	//	now = now.Next
	//}
	//for i := 0; i < n-k; i++ {
	//	head = head.Next
	//}
	//return head

	// 等距离双指针
	//former, latter := head, head
	//for i := 0; i < k; i++ {
	//	latter = latter.Next
	//}
	//for latter != nil {
	//	former = former.Next
	//	latter = latter.Next
	//}
	//return former

	// 等距离双指针，空间换时间的优化版本
	former, latter := head, head
	t := 0
	for latter != nil {
		if t >= k {
			former = former.Next
		}
		latter = latter.Next
		t++
	}
	return former
}

func main() {

}
