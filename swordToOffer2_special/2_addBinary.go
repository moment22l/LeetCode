package main

import "strconv"

// 	题目: 二进制加法
//	给定两个 01 字符串 a 和 b ，请计算它们的和，并以二进制字符串的形式输出。
//	输入为 非空 字符串且只包含数字 1 和 0。

func addBinary(a string, b string) string {
	if a == "0" {
		return b
	}
	if b == "0" {
		return a
	}
	if len(a) > len(b) {
		for i := len(a) - len(b); i > 0; i-- {
			b = "0" + b
		}
	} else if len(b) > len(a) {
		for i := len(b) - len(a); i > 0; i-- {
			a = "0" + a
		}
	}
	l := len(a)
	ans := ""
	var carry uint8 = 0
	for i := l - 1; i >= 0; i-- {
		posA := a[i] - 48
		posB := b[i] - 48
		add := posA + posB + carry
		carry = 0
		if add >= 2 {
			carry = 1
			add -= 2
		}
		addS := strconv.Itoa(int(add))
		ans = addS + ans
	}
	if carry == 1 {
		return "1" + ans
	}
	return ans
}
