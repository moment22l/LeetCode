package main

// 回溯 先加左括号, 然后判断是否可以加入右括号 时间O(4^n/(n*√n)) 空间O(2*n)=O(n)
func generateParenthesis(n int) (ans []string) {
	str := ""
	var backTrace func(left, right int)
	backTrace = func(left, right int) {
		if len(str) == n*2 {
			ans = append(ans, str)
			return
		}
		if left < n {
			str = str + "("
			backTrace(left+1, right)
			str = str[:len(str)-1]
		}
		if right < left && right < n {
			str = str + ")"
			backTrace(left, right+1)
			str = str[:len(str)-1]
		}
	}
	backTrace(0, 0)
	return
}
