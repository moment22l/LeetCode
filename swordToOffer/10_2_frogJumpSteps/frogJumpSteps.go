package main

import "fmt"

const mod int = 1e9 + 7

// var total int

func numWays(n int) int {
	// 回溯法
	// if n == 0 {
	// 	return 1
	// }
	// total = 0
	// traceBack(n, 0)
	// return total

	if n == 0 {
		return 1
	}
	a, b := 1, 1
	for i := 1; i < n; i++ {
		a, b = b, (a+b)%(1e9+7)
	}
	return b
}

// 回溯法，答案正确但超时
//func traceBack(n int, now int) {
//	if now >= n {
//		return
//	} else {
//		for i := 1; i <= 2; i++ {
//			if now+i == n {
//				total++
//				total = total % mod
//			}
//			traceBack(n, now+i)
//		}
//	}
//}

func main() {
	fmt.Println(numWays(7))
	fmt.Println(numWays(2))
	fmt.Println(numWays(0))
	fmt.Println(numWays(1))
	fmt.Println(numWays(41))
}
