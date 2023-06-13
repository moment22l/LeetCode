package main

import (
	"math"
)

// dp[i] = dp[j] && haveWord(s[j:i]), j = (0,1, ..., i-1)
func wordBreak(s string, wordDict []string) bool {
	headLetter := make(map[string]bool)
	minLen, maxLen := math.MaxInt, math.MinInt
	for _, word := range wordDict {
		headLetter[word] = true
		if len(word) < minLen {
			minLen = len(word)
		}
		if len(word) > maxLen {
			maxLen = len(word)
		}
	}
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := minLen; i <= n; i++ {
		for j := i - minLen; j >= i-maxLen && j >= 0; j-- {
			_, ok := headLetter[s[j:i]]
			if dp[j] && ok {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
