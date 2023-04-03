package main

type NumMatrix struct {
	sum [][]int // 记录每个位置到左上角的矩阵中的所有元素之和
}

func Constructor(matrix [][]int) NumMatrix {
	// 初始化二维数组
	m := NumMatrix{sum: make([][]int, len(matrix))}
	for i := 0; i < len(matrix); i++ {
		arr := make([]int, len(matrix[0]))
		m.sum[i] = arr
	}
	// 初始化最左列
	m.sum[0][0] = matrix[0][0]
	for i := 1; i < len(matrix); i++ {
		m.sum[i][0] = m.sum[i-1][0] + matrix[i][0]
	}
	// 初始化最上行
	for j := 1; j < len(matrix[0]); j++ {
		m.sum[0][j] = m.sum[0][j-1] + matrix[0][j]
	}
	// 动态规划计算每个从(0,0)到(i,j)矩阵的和
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			m.sum[i][j] = m.sum[i][j-1] + m.sum[i-1][j] - m.sum[i-1][j-1] + matrix[i][j]
		}
	}
	return m
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	// 注意判断row1和col1为0的情况
	if row1 == 0 && col1 == 0 {
		return this.sum[row2][col2]
	} else if row1 == 0 && col1 != 0 {
		return this.sum[row2][col2] - this.sum[row2][col1-1]
	} else if row1 != 0 && col1 == 0 {
		return this.sum[row2][col2] - this.sum[row1-1][col2]
	} else {
		return this.sum[row2][col2] - this.sum[row2][col1-1] - this.sum[row1-1][col2] + this.sum[row1-1][col1-1]
	}
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
