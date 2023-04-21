package main

// dp[j] = dp[j-1] + triangle[i][j], j == len(triangle[i])-1
// dp[j] = min{dp[j], dp[j-1]} + triangle[i][j], j != len(triangle[i])-1 && j != 0
// dp[j] = dp[j] + triangle[i][j], j == 0
// 时间O(mn) 空间O(n)
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	n := len(triangle[m-1])
	dp := make([]int, n)
	dp[0] = triangle[0][0]
	for i := 1; i < m; i++ {
		dp[i] = dp[i-1] + triangle[i][i]
		for j := i - 1; j > 0; j-- {
			dp[j] = min100(dp[j], dp[j-1]) + triangle[i][j]
		}
		dp[0] += triangle[i][0]
	}
	return minInSlice(dp)
}

func min100(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 不用这个函数能快一点, 因为调用这个函数会存在复制切片的开销, 但是逻辑更清晰一些
func minInSlice(arr []int) int {
	ans := arr[0]
	for _, num := range arr {
		if num < ans {
			ans = num
		}
	}
	return ans
}
