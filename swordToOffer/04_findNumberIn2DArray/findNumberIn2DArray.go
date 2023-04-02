package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	// 暴力解法
	// if len(matrix) == 0 || len(matrix[0]) == 0 {
	// 	return false
	// }
	// n := len(matrix)
	// m := len(matrix[0])
	// for i := 0; i < n; i++ {
	// 	for j := 0; j < m; j++ {
	// 		if matrix[i][j] == target {
	// 			return true
	// 		}
	// 		if matrix[i][j] > target {
	// 			break
	// 		}
	// 	}
	// }
	// return false

	// 稍微巧妙的解法
	// 思路(将整个数组逆时针转45°后类似于二叉搜索树):
	// 设big为第一个数组的最后一个数
	// (1) 从big开始向左遍历数组，直到找到第一个小于target的数small(由于非递减，所以该数右边的所有列都可以排除)
	// (2) 从small开始向下下遍历，直到找到一个大于target的数big(由于非递减，该数左上角的所有数均可以排除)
	// (3) 重复(1)(2)，直到找到target或索引越界
	//  [
	//  	[1,   4,  7, 11, 15],
	//  	[2,   5,  8, 12, 19],
	//  	[3,   6,  9, 16, 22],
	//  	[10, 13, 14, 17, 24],
	//  	[18, 21, 23, 26, 30]
	//  ]
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	n := len(matrix)
	m := len(matrix[0])
	i := 0
	j := m - 1
	for i < n {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
		if i > n-1 || j < 0 {
			break
		}
	}
	return false
}

func main() {

}
