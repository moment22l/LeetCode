package main

import "LeetCode/utils"

type CBTInserter struct {
	Root  *utils.TreeNode
	Queue []*utils.TreeNode
}

func Constructor_CBT(root *utils.TreeNode) CBTInserter {
	queue := []*utils.TreeNode{root}
	p := queue[0]
	for {
		if p.Left != nil {
			queue = append(queue, p.Left)
		} else {
			break
		}
		if p.Right != nil {
			queue = append(queue, p.Right)
		} else {
			break
		}
		queue = queue[1:]
		p = queue[0]
	}
	return CBTInserter{root, queue}
}

func (this *CBTInserter) Insert(v int) int {
	p := this.Queue[0]
	if p.Left == nil {
		p.Left = &utils.TreeNode{v, nil, nil}
		this.Queue = append(this.Queue, p.Left)
	} else {
		p.Right = &utils.TreeNode{v, nil, nil}
		this.Queue = append(this.Queue, p.Right)
		this.Queue = this.Queue[1:]
	}
	return p.Val
}

func (this *CBTInserter) Get_root() *utils.TreeNode {
	return this.Root
}
