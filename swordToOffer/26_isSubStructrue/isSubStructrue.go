package main

//
//import "fmt"
//
//// TreeNode Definition for a binary tree node.
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//// var nodeArr []*TreeNode
//
//func isSubStructure(A *TreeNode, B *TreeNode) bool {
//	if A == nil || B == nil {
//		return false
//	}
//	return findNode(A, B)
//}
//
//// 前序遍历寻找A中与B的根节点相等的节点, 并通过isEqual函数判断是否为True
//func findNode(now *TreeNode, B *TreeNode) bool {
//	if now.Val == B.Val {
//		if isEqual(now, B) {
//			return true
//		}
//	}
//	if now.Left != nil {
//		if findNode(now.Left, B) {
//			return true
//		}
//	}
//	if now.Right != nil {
//		if findNode(now.Right, B) {
//			return true
//		}
//	}
//	return false
//}
//
//// 按B的结构判断A为根节点是否能构造出B树
//func isEqual(A *TreeNode, B *TreeNode) bool {
//	if (A == nil && B != nil) || (A != nil && B == nil) || A.Val != B.Val {
//		return false
//	}
//	if B.Left != nil {
//		if !isEqual(A.Left, B.Left) {
//			return false
//		}
//	}
//	if B.Right != nil {
//		if !isEqual(A.Right, B.Right) {
//			return false
//		}
//	}
//	return true
//}
//
//func main() {
//	A := &TreeNode{4, &TreeNode{2, &TreeNode{4, &TreeNode{8, nil, nil}, &TreeNode{9, nil, nil}},
//		&TreeNode{5, nil, nil}}, &TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}}
//	B := &TreeNode{4, &TreeNode{8, nil, nil}, &TreeNode{9, nil, nil}}
//	fmt.Println(isSubStructure(A, B))
//}
