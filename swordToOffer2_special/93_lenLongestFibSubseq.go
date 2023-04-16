package main

// dp[i][j] 表示以第i个数为结尾, 第j个数为序列的倒数第二个数时序列的最长长度
// 时间O(n^2) 空间O(n^2)
func lenLongestFibSubseq(arr []int) (ans int) {
	n := len(arr)
	m := make(map[int]int, n)
	for i, num := range arr {
		m[num] = i
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, i)
	}
	for i := 1; i < n; i++ {
		dp[i][0] = 2
	}
	for i := 2; i < n; i++ {
		for j := 1; j < i; j++ {
			if k, ok := m[arr[i]-arr[j]]; ok && arr[i]-arr[j] < arr[j] {
				dp[i][j] = dp[j][k] + 1
			} else {
				dp[i][j] = 2
			}
			ans = max93(ans, dp[i][j])
		}
	}
	if ans == 2 {
		ans = 0
	}
	return
}

func max93(x, y int) int {
	if x > y {
		return x
	}
	return y
}
