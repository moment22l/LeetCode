package main

import "LeetCode/utils"

// flatten 展平多级双向链表
func flatten(root *utils.Node) *utils.Node {
	// 利用栈存储 每级中指向子链表节点的后一个节点, 时间O(n), 空间O(n)(ps: 空间的最坏情况为所有都有Child而没有next)
	stack := make([]*utils.Node, 0)
	p := root
	for p != nil {
		if p.Child != nil {
			if p.Next != nil {
				stack = append(stack, p.Next)
			}
			p.Next = p.Child
			p.Child.Prev = p
			p.Child = nil
		}
		if p.Next == nil {
			break
		}
		p = p.Next
	}
	for len(stack) != 0 {
		p.Next = stack[len(stack)-1]
		stack[len(stack)-1].Prev = p
		stack = stack[:len(stack)-1]
		for p.Next != nil {
			p = p.Next
		}
	}
	return root
}

// 另一种思路: 深度优先搜索, 遍历到有Child的节点就进入下一级, 直到最后遍历到nil再往前回溯
// 时间O(n), 空间O(n)(ps: 空间的最坏情况为所有都有Child而没有next)
