package main

import "fmt"

// Node Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	prev, cur := head, head
	for cur != nil {
		cur = cur.Next
		prev.Next = &Node{prev.Val, cur, nil}
		prev = prev.Next.Next
	}
	prev, cur = head, head.Next
	for cur != nil {
		if prev.Random != nil {
			cur.Random = prev.Random.Next
		}
		if cur.Next == nil {
			cur = cur.Next
		} else {
			cur = cur.Next.Next
		}
		prev = prev.Next.Next
	}
	copyHead := head.Next
	prev, cur = head, head.Next
	for cur != nil {
		prev.Next = cur.Next
		if cur.Next != nil {
			cur.Next = cur.Next.Next
		}
		prev = prev.Next
		cur = cur.Next
	}
	return copyHead
}

func main() {
	head := &Node{7, &Node{13, &Node{11, &Node{10, &Node{1, nil, nil}, nil}, nil}, nil}, nil}
	now := head
	now.Next.Random = now
	now.Next.Next.Next.Next.Random = now
	now = now.Next.Next
	now.Next.Random = now
	now.Random = now.Next.Next
	copyList := copyRandomList(head)
	fmt.Println(copyList)
}
