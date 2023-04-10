package main

import "LeetCode/utils"

//type BSTIterator struct {
//	queue []int
//}
//
//func ConstructorBSTIterator(root *utils.TreeNode) BSTIterator {
//	// 中序遍历将所有节点加入到队列中 时间O(n)
//	it := BSTIterator{queue: make([]int, 0)}
//	var inorder func(node *utils.TreeNode)
//	inorder = func(node *utils.TreeNode) {
//		if node == nil {
//			return
//		}
//		inorder(node.Left)
//		it.queue = append(it.queue, node.Val)
//		inorder(node.Right)
//	}
//	inorder(root)
//	return it
//}
//
//func (this *BSTIterator) Next() int {
//	// 每次从队列中取一个值 时间O(1)
//	ans := this.queue[0]
//	this.queue = this.queue[1:]
//	return ans
//}
//
//func (this *BSTIterator) HasNext() bool {
//	// 判断队列是否为空 时间O(1)
//	return len(this.queue) != 0
//}

type BSTIterator struct {
	// 空间O(n)
	stack []*utils.TreeNode
	cur   *utils.TreeNode
}

func ConstructorBSTIterator(root *utils.TreeNode) BSTIterator {
	// 时间O(1)
	return BSTIterator{cur: root}
}

func (this *BSTIterator) Next() int {
	// 单次调用最坏复杂度为O(n), 调用n次的均摊复杂度则为O(1)
	for node := this.cur; node != nil; node = node.Left {
		this.stack = append(this.stack, node)
	}
	this.cur, this.stack = this.stack[len(this.stack)-1], this.stack[:len(this.stack)-1]
	val := this.cur.Val
	this.cur = this.cur.Right
	return val
}

func (this *BSTIterator) HasNext() bool {
	// 栈为空且cur为空时才为false 时间O(1)
	return len(this.stack) != 0 || this.cur != nil
}
