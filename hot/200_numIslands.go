package main

func numIslands(grid [][]byte) (ans int) {
	m := len(grid)
	n := len(grid[0])
	alreadyPass := make([][]bool, m)
	for i := 0; i < m; i++ {
		alreadyPass[i] = make([]bool, n)
	}
	var pass func(int, int)
	pass = func(row, col int) {
		if grid[row][col] == '0' || alreadyPass[row][col] {
			return
		}
		alreadyPass[row][col] = true
		if col-1 >= 0 {
			pass(row, col-1)
		}
		if col+1 < n {
			pass(row, col+1)
		}
		if row-1 >= 0 {
			pass(row-1, col)
		}
		if row+1 < m {
			pass(row+1, col)
		}
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' && !alreadyPass[i][j] {
				pass(i, j)
				ans++
			}
		}
	}
	return ans
}
