package main

import "LeetCode/utils"

func findBottomLeftValue(root *utils.TreeNode) (ans int) {
	queue := make([]*utils.TreeNode, 1)
	queue[0] = root
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		ans = node.Val
	}
	return
}
