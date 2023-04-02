package main

import "fmt"

// 回溯
var combinations []string

func generateParenthesis(n int) []string {
	combinations = []string{}
	backTrack(n, 0, 0, []byte{})
	return combinations
}

func backTrack(n int, left int, right int, stack []byte) {
	if len(stack) == n*2 {
		combinations = append(combinations, string(stack))
	}
	if left < n {
		stack = append(stack, '(')
		backTrack(n, left+1, right, stack)
		stack = stack[0 : len(stack)-1]
	}
	if right < left {
		stack = append(stack, ')')
		backTrack(n, left, right+1, stack)
		stack = stack[0 : len(stack)-1]
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(1))
}
