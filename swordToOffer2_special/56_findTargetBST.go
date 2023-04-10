package main

import "LeetCode/utils"

func findTarget(root *utils.TreeNode, k int) bool {
	// 比较暴力的解法 但简单易懂而且不慢 时间O(n) 空间O(n)
	// 中序遍历BST 将所有节点值存入数组中, 然后用双指针找到和为k的数(也可将数组换成哈希表, 那就可以在递归的时候就得出答案)
	arr := make([]int, 0)
	var inorder func(node *utils.TreeNode)
	inorder = func(node *utils.TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		arr = append(arr, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	for i, j := 0, len(arr)-1; i < j; {
		total := arr[i] + arr[j]
		if total == k {
			return true
		} else if total > k {
			j--
		} else {
			i++
		}
	}
	return false
}
