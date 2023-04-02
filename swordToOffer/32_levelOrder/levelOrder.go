package main

//// TreeNode Definition for a binary tree node.
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var queue []*TreeNode
	var temp *TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		res = append(res, queue[0].Val)
		temp = queue[0]
		if temp.Left != nil {
			queue = append(queue, temp.Left)
		}
		if temp.Right != nil {
			queue = append(queue, temp.Right)
		}
		queue = queue[1:]
	}
	return res
}
