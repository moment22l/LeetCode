package main

import "fmt"

// 通过回溯遍历所有可能情况
var letterMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		digit := string(digits[index])
		letters := letterMap[digit]
		for i := 0; i < len(letters); i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
	digits = ""
	fmt.Println(letterCombinations(digits))
	digits = "2"
	fmt.Println(letterCombinations(digits))
	digits = "23456789"
	fmt.Println(letterCombinations(digits))
}
