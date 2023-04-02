package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return (A != nil && B != nil) && (findTree(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}

func findTree(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return findTree(A.Left, B.Left) && findTree(A.Right, B.Right)
}

func main() {

}
