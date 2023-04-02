package main

import (
	"LeetCode/swordToOffer/utils/tree"
	"fmt"
)

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pathCol [][]int

func pathSum(root *tree.TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}
	pathCol = [][]int{}
	dfs(root, target, []int{})
	return pathCol
}

func dfs(node *tree.TreeNode, target int, path []int) {
	if node.Val == target && node.Left == nil && node.Right == nil {
		path = append(path, node.Val)
		temp := make([]int, len(path))
		copy(temp, path)
		pathCol = append(pathCol, temp)
	} else {
		if node.Left != nil {
			dfs(node.Left, target-node.Val, append(path, node.Val))
		}
		if node.Right != nil {
			dfs(node.Right, target-node.Val, append(path, node.Val))
		}
	}
}

func main() {
	tree := tree.BinaryTreeGenerator("C:\\Users\\82507\\AppData\\Roaming\\JetBrains" +
		"\\IntelliJIdea2022.2\\scratches\\scratch_1.json")
	fmt.Println(pathSum(tree, -243))
}
