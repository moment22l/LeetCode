package main

func findTheLongestBalancedSubstring(s string) (ans int) {
	count0, count1 := 0, 0
	for _, ch := range s {
		if ch == '0' {
			if count1 != 0 {
				count0, count1 = 0, 0
			}
			count0++
		} else {
			count1++
			ans = max2609(ans, min2609(count0, count1)*2)
		}
	}
	return
}

func max2609(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min2609(x, y int) int {
	if x < y {
		return x
	}
	return y
}
