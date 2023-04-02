package main

import (
	"bytes"
	"fmt"
)

func isNumber(s string) bool {
	bs := []byte(s)
	bs = bytes.TrimSpace(bs)
	state := make([]map[byte]int, len(bs))
	// ' '表示当前字符为空格, 's'表示当前字符为正负号, 'n'表示当前字符为数字
	// 'p'表示当前字符为小数点, 'e'表示当前字符为幂符号
	//state = []map[byte]int{
	//	{' ': 0, 's': 1, 'n': 2, 'p': 3},
	//	{'n': 2, 'p': 3},
	//	{'n': 2, 'p': 4, 'e': 5, ' ': 8},
	//	{'n': 4},
	//	{'n': 4, 'e': 5, ' ': 8},
	//	{'s': 6, 'n': 7},
	//	{'n': 7},
	//	{'n': 7, ' ': 8},
	//	{' ': 8},
	//}
	state = []map[byte]int{
		{'s': 1, 'n': 2, 'p': 3},
		{'n': 2, 'p': 3},
		{'n': 2, 'p': 4, 'e': 5},
		{'n': 4},
		{'n': 4, 'e': 5},
		{'s': 6, 'n': 7},
		{'n': 7},
		{'n': 7},
	}
	var c byte
	nowState := 0
	for _, b := range bs {
		if b >= '0' && b <= '9' {
			c = 'n'
		} else if b == '+' || b == '-' {
			c = 's'
		} else if b == 'e' || b == 'E' {
			c = 'e'
		} else if b == '.' {
			c = 'p'
		} else {
			return false
		}
		if _, ok := state[nowState][c]; !ok {
			return false
		}
		nowState = state[nowState][c]
	}
	if nowState == 2 || nowState == 4 || nowState == 7 || nowState == 8 {
		return true
	}
	return false
}

func main() {
	s := ".1"
	fmt.Println(isNumber(s))
}
