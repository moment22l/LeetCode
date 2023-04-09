package main

import "LeetCode/utils"

func pruneTree(root *utils.TreeNode) *utils.TreeNode {
	// 后序遍历逐步删除
	if root == nil {
		return root
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return root
}
