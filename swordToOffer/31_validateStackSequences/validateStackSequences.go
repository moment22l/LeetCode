package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	// 直接模拟真实入栈出栈操作，观察最后是否能将所有元素都按popped顺序入栈并出栈
	var stack []int
	popIndex := 0
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && popped[popIndex] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			popIndex++
		}
	}
	return len(stack) == 0
}

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 3, 5, 1, 2}
	fmt.Println(validateStackSequences(pushed, popped))
}
