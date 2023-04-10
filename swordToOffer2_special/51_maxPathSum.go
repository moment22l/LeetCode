package main

import (
	"LeetCode/utils"
	"math"
)

func maxPathSum(root *utils.TreeNode) int {
	// 递归计算每个节点的贡献值(以该节点为起点的最大路径和)
	// 然后逐步判断max
	// 时间O(n) 空间O(n)
	maxSum := math.MinInt
	var maxGain func(node *utils.TreeNode) int
	maxGain = func(node *utils.TreeNode) int {
		// 后序遍历
		if node == nil {
			return 0
		}
		// 贡献值大于0才会加入到路径中
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)
		// 计算当前最大路径之和并与当前最佳路径和比较
		pathTotal := node.Val + leftGain + rightGain
		maxSum = max(pathTotal, maxSum)
		// 返回当前节点的最大贡献值
		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
