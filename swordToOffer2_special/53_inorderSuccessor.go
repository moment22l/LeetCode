package main

import "LeetCode/utils"

//func inorderSuccessor(root *utils.TreeNode, p *utils.TreeNode) *utils.TreeNode {
//	// 递归中序遍历并用标志代表是否已经找到  缺点：仍然会遍历整棵树  优化：用递归找到就返回
//	// 时间O(n) 空间O(n)
//	flag := false
//	var ans *utils.TreeNode
//	var inorderFind func(*utils.TreeNode)
//	inorderFind = func(node *utils.TreeNode) {
//		if node == nil {
//			return
//		}
//		inorderFind(node.Left)
//		if flag {
//			ans = node
//			flag = false
//			return
//		}
//		if node == p {
//			flag = true
//		}
//		inorderFind(node.Right)
//	}
//	inorderFind(root)
//	return ans
//}

//func inorderSuccessor(root *utils.TreeNode, p *utils.TreeNode) *utils.TreeNode {
//	// 优化：用递归中续遍历找到就可以直接返回
//	// 时间O(n) 空间O(n)
//	var s []*utils.TreeNode
//	var pre, cur *utils.TreeNode = nil, root
//	for len(s) != 0 || cur != nil {
//		for cur != nil {
//			s = append(s, cur)
//			cur = cur.Left
//		}
//		cur = s[len(s)-1]
//		s = s[:len(s)-1]
//		if pre == p {
//			return cur
//		}
//		pre = cur
//		cur = cur.Right
//	}
//	return nil
//}

func inorderSuccessor(root *utils.TreeNode, p *utils.TreeNode) *utils.TreeNode {
	// 利用二叉搜索树的性质
	// 时间O(n) 空间O(1)

	// p.Right不为nil的时候, 结果一定是p.Right为根的树中最左边那个节点
	var ans *utils.TreeNode
	if p.Right != nil {
		ans = p.Right
		for ans.Left != nil {
			ans = ans.Left
		}
		return ans
	}
	// p.Right为nil的时候, 遍历找p的祖先节点
	// 当前节点比p大时, 更新ans并让node往左走
	// 节点比p小或等于p时, 不更新ans并让node往右走
	node := root
	for node != nil {
		if node.Val > p.Val {
			ans = node
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return ans
}
