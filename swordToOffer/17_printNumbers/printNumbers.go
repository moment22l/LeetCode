package main

import (
	"fmt"
	"strconv"
)

var numbers []string
var num string = ""

func printNumbers(n int) []string {
	// 不考虑大数打印
	//var numbers []int
	//for i := 1; i < int(math.Pow10(n)); i++ {
	//	numbers = append(numbers, i)
	//}
	//return numbers

	// 抛开返回[]int的限制，考虑大数打印
	numbers = []string{}
	if n > 0 {
		for i := 1; i <= n; i++ {
			dfs(i, 0)
		}
	}
	return numbers
}

func dfs(n int, now int) {
	if now == n {
		// temp, _ := strconv.Atoi(num)
		numbers = append(numbers, num)
		return
	}
	start := 0
	if now == 0 {
		start++
	}
	for i := start; i <= 9; i++ {
		num = num + strconv.Itoa(i)
		dfs(n, now+1)
		num = num[:len(num)-1]
	}
}

func main() {
	printNumbers(4)
	fmt.Println(numbers)
}
