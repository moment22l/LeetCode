package main

// 回溯
var combinations22 []string

func generateParenthesis(n int) []string {
	combinations22 = []string{}
	backTrack22(n, 0, 0, []byte{})
	return combinations22
}

func backTrack22(n int, left int, right int, stack []byte) {
	if len(stack) == n*2 {
		combinations22 = append(combinations22, string(stack))
	}
	if left < n {
		stack = append(stack, '(')
		backTrack22(n, left+1, right, stack)
		stack = stack[0 : len(stack)-1]
	}
	if right < left {
		stack = append(stack, ')')
		backTrack22(n, left, right+1, stack)
		stack = stack[0 : len(stack)-1]
	}
}
