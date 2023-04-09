package main

import "LeetCode/utils"

// 时间O(n) 空间O(n)(递归需要堆栈空间)
func sumNumbers(root *utils.TreeNode) int {
	ans := 0
	now := 0
	var dfs func(*utils.TreeNode)
	dfs = func(root *utils.TreeNode) {
		if root == nil {
			return
		}
		now = now*10 + root.Val
		dfs(root.Left)
		dfs(root.Right)
		if root.Left == nil && root.Right == nil {
			ans += now
		}
		now = now / 10
	}
	dfs(root)
	return ans
}
