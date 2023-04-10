package main

import "LeetCode/utils"

// 递归 时间O(n) 空间O(n)
func increasingBST(root *utils.TreeNode) *utils.TreeNode {
	pre := &utils.TreeNode{}
	res := pre
	var inorder func(*utils.TreeNode)
	inorder = func(node *utils.TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res.Right = node
		node.Left = nil
		res = node
		inorder(node.Right)
	}
	inorder(root)
	return root
}

// 迭代 中序遍历过程中更改链表 时间O(n) 空间O(n)
//func increasingBST(root *utils.TreeNode) *utils.TreeNode {
//	stack := make([]*utils.TreeNode, 1)
//	stack[0] = root
//	flag := false
//	var inorder func(*utils.TreeNode)
//	inorder = func(*utils.TreeNode) {
//		var node *utils.TreeNode
//		for len(stack) != 0 {
//			if stack[len(stack)-1].Left != nil {
//				stack = append(stack, stack[len(stack)-1].Left)
//				continue
//			}
//			node = stack[len(stack)-1]
//			if !flag {
//				flag = true
//				root = node
//			}
//			stack = stack[:len(stack)-1]
//			if node.Right != nil {
//				stack = append(stack, node.Right)
//				for stack[len(stack)-1].Left != nil {
//					stack = append(stack, stack[len(stack)-1].Left)
//				}
//			}
//			if len(stack) != 0 {
//				node.Right = stack[len(stack)-1]
//				stack[len(stack)-1].Left = nil
//			}
//		}
//	}
//	inorder(root)
//	return root
//}
