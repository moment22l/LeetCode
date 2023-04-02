package main

import "fmt"

// TreeNode Definition for a binary tree node.
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func levelOrder_2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	queue := []*TreeNode{root}
	var temp *TreeNode
	var tmp []int
	for len(queue) != 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			temp = queue[0]
			queue = queue[1:]
			tmp = append(tmp, temp.Val)
			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
		}
		res = append(res, tmp)
		tmp = []int{}
	}
	return res
}

func main() {
	fmt.Println(levelOrder_2(&TreeNode{1, &TreeNode{2, nil, nil}, nil}))
}
