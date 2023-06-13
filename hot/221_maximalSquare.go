package main

// dp[i][j]表示以(i, j)为右下角的最大正方形的边长
// dp[i][j]的情况:
// 判断matrix[i][j]是否为1, 若不为1则 dp[i][j]=0
// 若为1则:
// 判断matrix[i-1][j-1]是否为1, 若不为1则 dp[i][j] = matrix[i][j]
// 若为1则:
// 依次判断(i,j)位置的左边和上边的元素是否为1, 若不为1则 终止
// 若为1则:
// dp[i][j]++
func maximalSquare(matrix [][]byte) int {
	// 初始化dp
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		dp[i][0] = int(matrix[i][0] - '0')
	}
	for j := 0; j < n; j++ {
		dp[0][j] = int(matrix[0][j] - '0')
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = min221_3(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
		}
	}

	maxLen := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] > maxLen {
				maxLen = dp[i][j]
			}
		}
	}
	return maxLen * maxLen
}

func min221_3(i int, i2 int, i3 int) int {
	ans := i
	if ans > i2 {
		ans = i2
	}
	if ans > i3 {
		ans = i3
	}
	return ans
}
