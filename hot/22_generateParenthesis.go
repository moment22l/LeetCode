package main

// 暴力递归
//var combinations []string
//
//func generateParenthesis(n int) []string {
//	combinations = []string{}
//	backTrack(n, 0, "")
//	return combinations
//}
//
//func backTrack(n int, index int, combination string) {
//	if index == n*2 {
//		if isValid(combination) {
//			combinations = append(combinations, combination)
//		}
//	} else {
//		for i := 0; i < 2; i++ {
//			backTrack(n, index+1, combination+string(byte(i+40)))
//		}
//	}
//}
//
//func isValid(s string) bool {
//	stack := make([]byte, 0)
//	for _, b := range s {
//		if b == '(' {
//			stack = append(stack, byte(b))
//		} else {
//			if len(stack) == 0 {
//				return false
//			}
//			stack = stack[0 : len(stack)-1]
//		}
//	}
//	if len(stack) != 0 {
//		return false
//	}
//	return true
//}
