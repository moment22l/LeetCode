package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	res := make([]int, len(matrix)*len(matrix[0]))
	index := 0
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			res[index] = matrix[top][i]
			index++
		}
		top++
		if top > bottom {
			break
		}
		for i := top; i <= bottom; i++ {
			res[index] = matrix[i][right]
			index++
		}
		right--
		if right < left {
			break
		}
		for i := right; i >= left; i-- {
			res[index] = matrix[bottom][i]
			index++
		}
		bottom--
		if bottom < top {
			break
		}
		for i := bottom; i >= top; i-- {
			res[index] = matrix[i][left]
			index++
		}
		left++
		if left > right {
			break
		}
	}
	return res
}

func main() {
	matrix := [][]int{[]int{1, 2, 3}}
	fmt.Println(spiralOrder(matrix))
}
