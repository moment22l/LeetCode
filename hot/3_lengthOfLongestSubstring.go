package main

import (
	"math"
)

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	lastOccuredMap := make(map[rune]int)
	start := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccuredMap[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		maxLen = int(math.Max(float64(i-start+1), float64(maxLen)))
		lastOccuredMap[ch] = i
	}
	return maxLen
}
