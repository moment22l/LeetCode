package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder_3(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	queue := []*TreeNode{root}
	var temp *TreeNode
	for len(queue) != 0 {
		n := len(queue)
		tmp := make([]int, n)
		for i := 0; i < n; i++ {
			temp = queue[0]
			queue = queue[1:]
			if len(res)%2 == 0 {
				tmp[i] = temp.Val
			} else {
				tmp[n-1-i] = temp.Val
			}
			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
		}
		res = append(res, tmp)
	}
	return res
}
