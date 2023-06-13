package main

func addStrings(num1 string, num2 string) (res string) {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	n1, n2 := len(num1), len(num2)
	var carry byte
	i, j := n1-1, n2-1
	for ; j >= 0; i, j = i-1, j-1 {
		now := (num1[i] - '0') + (num2[j] - '0') + carry
		carry = now / 10
		res = string(now-carry*10+'0') + res
	}
	for ; carry != 0 && i >= 0; i-- {
		now := (num1[i] - '0') + carry
		carry = now / 10
		res = string(now-carry*10+'0') + res
	}
	if i < 0 && carry != 0 {
		res = string(carry+'0') + res
	} else if i >= 0 {
		res = num1[:i+1] + res
	}
	return res
}
