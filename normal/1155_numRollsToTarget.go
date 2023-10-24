package main

func numRollsToTarget(n int, k int, target int) int {
	mod := int(1e9 + 7)
	counts := make([]int, target+1)
	counts[0] = 1
	for i := 1; i <= n; i++ {
		for j := target; j >= 0; j-- {
			counts[j] = 0
			for x := 1; j-x >= 0 && x <= k; x++ {
				counts[j] = (counts[j] + counts[j-x]) % mod
			}
		}
	}
	return counts[target]
}
