package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return judge(root, root)
}

func judge(root, copyRoot *TreeNode) bool {
	if root == nil && copyRoot == nil {
		return true
	}
	if root == nil || copyRoot == nil {
		return false
	}
	return root.Val == copyRoot.Val && judge(root.Left, copyRoot.Right) && judge(root.Right, copyRoot.Left)
}

func main() {

}
