package main

// ListNode 链表节点 使用到的题目: 21-27
type ListNode struct {
	Val  int
	Next *ListNode
}

// Node 多级双向链表节点 使用到的题目: 28
type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}