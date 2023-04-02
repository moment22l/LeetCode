package main

import "fmt"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var val int
var copyK int

func kthLargest(root *TreeNode, k int) int {
	copyK = k
	midTraverse(root)
	return val
}

func midTraverse(root *TreeNode) {
	// 中序倒序遍历到第K个节点时返回
	if root == nil {
		return
	}
	midTraverse(root.Right)
	if copyK == 0 {
		return
	}
	copyK--
	if copyK == 0 {
		val = root.Val
	}
	midTraverse(root.Left)
}

func main() {
	tree := &TreeNode{3, &TreeNode{1, nil, &TreeNode{2, nil, nil}}, &TreeNode{4, nil, nil}}
	fmt.Println(kthLargest(tree, 1))
}
