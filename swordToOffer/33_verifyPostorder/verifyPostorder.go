package verifyPostorder

import (
	"math"
)

func verifyPostorder(postorder []int) bool {
	// 递归分治
	//return recur(postorder, 0, len(postorder)-1)

	if len(postorder) == 0 {
		return true
	}
	// 逆转数组
	for i := 0; i < len(postorder)/2; i++ {
		temp := postorder[i]
		postorder[i] = postorder[len(postorder)-1-i]
		postorder[len(postorder)-1-i] = temp
	}
	stack := make([]int, 0)
	root := math.MaxInt
	for i := 0; i < len(postorder); i++ {
		if postorder[i] > root {
			return false
		}
		for len(stack) != 0 {
			if postorder[i] > stack[len(stack)-1] {
				break
			}
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, postorder[i])
	}
	return true
}

//func recur(postorder []int, i, j int) bool {
//	if i >= j {
//		return true
//	}
//	n := i
//	for ; n < j; n++ {
//		if postorder[n] > postorder[j] {
//			break
//		}
//	}
//	p := n
//	for ; p < j; p++ {
//		if postorder[p] < postorder[j] {
//			break
//		}
//	}
//	return p == j && recur(postorder, i, n-1) && recur(postorder, n, j-1)
//}
