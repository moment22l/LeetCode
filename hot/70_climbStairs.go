package main

// 斐波那契数列
// 时间O(n) 空间O(1)
//func climbStairs(n int) int {
//	pre, cur := 0, 1
//	for i := 0; i < n; i++ {
//		pre, cur = cur, pre+cur
//	}
//	return cur
//}

// 数学算法 矩阵快速幂
type matrix [2][2]int

func climbStairs(n int) int {
	ans := pow(matrix{[2]int{1, 1}, [2]int{1, 0}}, n)
	return ans[1][0] + ans[1][1]
}

func mul(a, b matrix) (c matrix) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = a[i][0]*b[0][j] + a[i][1]*b[1][j]
		}
	}
	return
}

func pow(a matrix, n int) matrix {
	res := matrix{[2]int{1, 0}, [2]int{0, 1}}
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = mul(res, a)
		}
		a = mul(a, a)
	}
	return res
}
