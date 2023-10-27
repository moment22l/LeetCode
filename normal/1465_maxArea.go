package main

import (
	"math"
	"sort"
)

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	return (calMax(horizontalCuts, h) * calMax(verticalCuts, w)) % (int(math.Pow10(9)) + 7)
}

func calMax(arr []int, brd int) int {
	pre, ans := 0, 0
	for _, now := range arr {
		ans = max_1465(now-pre, ans)
		pre = now
	}
	ans = max_1465(ans, brd-pre)
	return ans
}

func max_1465(x, y int) int {
	if x > y {
		return x
	}
	return y
}
