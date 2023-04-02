package main

import "fmt"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// 递归
	// if len(preorder) == 0 {
	// 	return nil
	// }
	// root := &TreeNode{preorder[0], nil, nil}
	// i := 0
	// for ; i < len(inorder); i++ {
	// 	if inorder[i] == preorder[0] {
	// 		break
	// 	}
	// }
	// root.Left = buildTree(preorder[1:i+1], inorder[:i])
	// root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	// return root

	// 迭代
	// 1. 一直向左子树遍历，直到当前节点没有左儿子(即node.Val == inorder[inorderIndex])，
	//    此时就已建好根节点的最左边一个链表，且链表顺序与栈顺序相同(PS:inorderIndex起始值为0)
	// 2. 从栈顶开始向下寻找，同时令inorderIndex逐步增加，直到栈顶元素不等于inorder[inorderIndex]
	//    然后让inorder[inorderIndex]入栈，并令最后一个从栈中弹出的节点的右儿子为inorder[inorderIndex]
	// 3. 重复1.2, 直到preorder所有节点均已遍历结束
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	var stack []*TreeNode
	stack = append(stack, root)
	inorderIndex := 0
	for i := 1; i < len(preorder); i++ {
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{preorder[i], nil, nil}
			stack = append(stack, node.Left)
		} else {
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			node.Right = &TreeNode{preorder[i], nil, nil}
			stack = append(stack, node.Right)
		}
	}
	return root
}

func main() {
	temp := []int{9}
	fmt.Println(temp[0:1])
}
