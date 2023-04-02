package main

import "fmt"

func firstUniqChar(s string) byte {
	m := [26]int{}
	for _, b := range s {
		m[b-'a']++
	}
	for _, b := range s {
		if m[b-'a'] == 1 {
			return byte(b)
		}
	}
	return ' '
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("cc"))
}
