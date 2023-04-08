package main

type CBTInserter struct {
	Root  *TreeNode
	Queue []*TreeNode
}

func Constructor_CBT(root *TreeNode) CBTInserter {
	queue := []*TreeNode{root}
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
		p.Left = &TreeNode{v, nil, nil}
		this.Queue = append(this.Queue, p.Left)
	} else {
		p.Right = &TreeNode{v, nil, nil}
		this.Queue = append(this.Queue, p.Right)
		this.Queue = this.Queue[1:]
	}
	return p.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.Root
}
