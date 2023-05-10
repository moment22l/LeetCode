package main

import "LeetCode/utils"

// 递归完成中序遍历
func inorderTraversal(root *utils.TreeNode) (res []int) {
	var recursion func(node *utils.TreeNode)
	recursion = func(node *utils.TreeNode) {
		if node == nil {
			return
		}
		recursion(node.Left)
		res = append(res, node.Val)
		recursion(node.Right)
	}
	recursion(root)
	return
}
