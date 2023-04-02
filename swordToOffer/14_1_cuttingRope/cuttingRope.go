package main

import "math"

func cuttingRope(n int) int {
	// 动态规划
	// dp := make([]int, n+1)
	// for i := 2; i <= n; i++ {
	// 	nowMax := 0
	// 	for j := 1; j < i; j++ {
	// 		nowMax = max(nowMax, max(j*(i-j), j*dp[i-j]))
	// 	}
	// 	dp[i] = nowMax
	// }
	// return dp[n]

	// 数学推导
	// 推论一: 若把绳子分成a段，则当所有段都相等时，乘积最大
	// 推论二: 若按长度x分成a段，则当x=3时，乘积最大
	// 那么将绳子尽可能分成长度为3(其次为2)的段，则可以让乘积达到最大
	if n <= 3 {
		return n - 1
	}
	a, b := n/3, n%3
	if b == 0 {
		return int(math.Pow(3, float64(a)))
	} else if b == 1 {
		return int(math.Pow(3, float64(a-1))) * 4
	}
	return int(math.Pow(3, float64(a))) * 2
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {

}
