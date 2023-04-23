package main

func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}
	// 建立一个表存储每个括号对应的另一个括号
	var bracketsMap map[byte]byte = map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := make([]byte, 0) //
	for _, b := range s {
		if b == '(' || b == '{' || b == '[' {
			stack = append(stack, byte(b))
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			now := byte(b)
			if bracketsMap[now] != top {
				return false
			}
			stack = stack[0 : len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
