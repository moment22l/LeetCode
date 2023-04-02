package main

import "fmt"

func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n)
	dp[0] = 1
	a, b, c := 0, 0, 0
	for i := 1; i < n; i++ {
		dp[i] = Min(dp, a, b, c)
		if dp[a]*2 == dp[i] {
			a++
		}
		if dp[b]*3 == dp[i] {
			b++
		}
		if dp[c]*5 == dp[i] {
			c++
		}
	}
	return dp[n-1]
}

func Min(dp []int, a, b, c int) int {
	min := dp[a] * 2
	if dp[b]*3 < min {
		min = dp[b] * 3
	}
	if dp[c]*5 < min {
		min = dp[c] * 5
	}
	return min
}

func main() {
	fmt.Println(nthUglyNumber(10))
}
