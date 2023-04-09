package main

import "LeetCode/utils"

func largestValues(root *utils.TreeNode) (ans []int) {
	if root == nil {
		ans = []int{}
		return
	}
	queue := make([]*utils.TreeNode, 1)
	queue[0] = root
	for len(queue) != 0 {
		max := queue[0].Val
		index := len(queue)
		for i := 0; i < index; i++ {
			node := queue[i]
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, max)
		queue = queue[index:]
	}
	return
}
