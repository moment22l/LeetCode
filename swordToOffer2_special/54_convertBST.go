package main

import "LeetCode/utils"

//func convertBST(root *utils.TreeNode) *utils.TreeNode {
//  // 中序逆序遍历, 从大到小累加并赋值给每个节点 时间O(n) 空间O(n)
//	now := 0
//	var reverseInorder func(*utils.TreeNode)
//	reverseInorder = func(node *utils.TreeNode) {
//		if node == nil {
//			return
//		}
//		reverseInorder(node.Right)
//		now += node.Val
//		node.Val = now
//		reverseInorder(node.Left)
//	}
//	reverseInorder(root)
//	return root
//}

func getSuccessor(node *utils.TreeNode) *utils.TreeNode {
	succ := node.Right
	for succ.Left != nil && succ.Left != node {
		succ = succ.Left
	}
	return succ
}

func convertBST(root *utils.TreeNode) *utils.TreeNode {
	// Morris 遍历 时间O(n) 空间O(1) 太巧妙了
	sum := 0
	node := root
	for node != nil {
		if node.Right == nil {
			sum += node.Val
			node.Val = sum
			node = node.Left
		} else {
			succ := getSuccessor(node)
			if succ.Left == nil {
				succ.Left = node
				node = node.Right
			} else {
				succ.Left = nil
				sum += node.Val
				node.Val = sum
				node = node.Left
			}
		}
	}
	return root
}
