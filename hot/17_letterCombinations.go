package main

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

var combinations17 []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations17 = []string{}
	backtrack(digits, 0, "")
	return combinations17
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations17 = append(combinations17, combination)
	} else {
		digit := string(digits[index])
		letters := letterMap[digit]
		for i := 0; i < len(letters); i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}
