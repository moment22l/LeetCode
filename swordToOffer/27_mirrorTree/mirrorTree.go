package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	mirror(root)
	return root
}

func mirror(root *TreeNode) {
	if root.Left != nil {
		mirror(root.Left)
	}
	if root.Right != nil {
		mirror(root.Right)
	}
	swap(root)
}

func swap(node *TreeNode) {
	temp := node.Left
	node.Left = node.Right
	node.Right = temp
}

func main() {

}
