package main

func lengthOfLongestSubstring(s string) int {
	lastOccured := make(map[rune]int)
	maxLen := 0
	i := -1
	for j, ch := range s {
		if lastI, ok := lastOccured[ch]; ok && lastI > i {
			i = lastI
		} else {
			maxLen = max(maxLen, j-i)
		}
		lastOccured[ch] = j
	}
	return maxLen
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {

}
