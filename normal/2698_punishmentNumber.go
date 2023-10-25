package main

import "strconv"

func punishmentNumber(n int) (ans int) {
	var dfs func(string, int, int, int) bool
	dfs = func(num string, pos int, total int, target int) bool {
		if pos == len(num) {
			return total == target
		}
		sum := 0
		for i := pos; i < len(num); i++ {
			sum = sum*10 + int(num[i]-'0')
			if sum+total > target {
				break
			}
			if dfs(num, i+1, total+sum, target) {
				return true
			}
		}
		return false
	}
	for i := 1; i <= n; i++ {
		if dfs(strconv.Itoa(i*i), 0, 0, i) {
			ans += i * i
		}
	}
	return
}
