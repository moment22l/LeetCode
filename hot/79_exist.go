package main

type pair struct {
	x, y int
}

var directions = []pair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

// 深度优先搜索 回溯
// O(MN⋅(3^L)) 空间O(MN)
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		vis[i][j] = true
		defer func() { vis[i][j] = false }()
		for _, dir := range directions {
			newI, newJ := i+dir.x, j+dir.y
			if 0 <= newI && newI < m && 0 <= newJ && newJ < n && !vis[newI][newJ] {
				if dfs(newI, newJ, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
