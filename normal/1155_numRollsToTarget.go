package main

func numRollsToTarget(n int, k int, target int) int {
	// 动态规划dp的方式
	//mod := int(1e9 + 7)
	//counts := make([]int, target+1)
	//counts[0] = 1
	//for i := 1; i <= n; i++ {
	//	for j := target; j >= 0; j-- {
	//		counts[j] = 0
	//		for x := 1; j-x >= 0 && x <= k; x++ {
	//			counts[j] = (counts[j] + counts[j-x]) % mod
	//		}
	//	}
	//}
	//return counts[target]

	// 记忆化深度搜索 = 递归＋记录返回值
	if target < n || target > n*k {
		return 0 // 无法组成 target
	}
	const mod = 1_000_000_007
	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, target-n+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == 0 {
			if j == 0 {
				return 1
			}
			return 0
		}
		p := &memo[i][j]
		if *p != -1 { // 之前计算过
			return *p
		}
		res := 0
		for x := 0; x < k && x <= j; x++ { // 掷出了 x
			res = (res + dfs(i-1, j-x)) % mod
		}
		*p = res // 记忆化
		return res
	}
	return dfs(n, target-n)
}
