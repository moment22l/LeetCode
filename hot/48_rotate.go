package main

// 时间O(n^2) 空间O(1)
func rotate(matrix [][]int) {
	// 根据公式得某个转换公式
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i] =
				matrix[n-j-1][i], matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1]
		}
	}
	// 转置矩阵
	//for i := 0; i < len(matrix); i++ {
	//	for j := i; j < len(matrix); j++ {
	//		matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
	//	}
	//}
	//// 左右翻转
	//for i := 0; i < len(matrix); i++ {
	//	for j := 0; j < len(matrix)/2; j++ {
	//		matrix[i][j], matrix[i][len(matrix)-j-1] = matrix[i][len(matrix)-j-1], matrix[i][j]
	//	}
	//}
	return
}
