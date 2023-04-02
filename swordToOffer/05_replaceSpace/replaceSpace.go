package main

import "fmt"

func replaceSpace(s string) string {
	var newStr []byte
	for _, b := range []byte(s) {
		if b == ' ' {
			newStr = append(newStr, '%')
			newStr = append(newStr, '2')
			newStr = append(newStr, '0')
		} else {
			newStr = append(newStr, b)
		}
	}
	return string(newStr)
}

func main() {
	s := "We are happy."
	fmt.Println(replaceSpace(s))
}
