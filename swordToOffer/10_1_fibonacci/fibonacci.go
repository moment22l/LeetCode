package main

import "fmt"

const mod int = 1e9 + 7

type matrix [2][2]int

func multiply(a, b matrix) (c matrix) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = (a[i][0]*b[0][j] + a[i][1]*b[1][j]) % mod
		}
	}
	return
}

// 矩阵快速幂
func pow(a matrix, n int) matrix {
	ret := matrix{{1, 0}, {0, 1}}
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			ret = multiply(ret, a)
		}
		a = multiply(a, a)
	}
	return ret
}

func fib(n int) int {
	// 方法一: 建立长度为n的斐波那契数列数组，取最后一个
	// const mod int = 1e9 + 7
	// if n < 2 {
	// 	return n
	// }
	// nums := make([]int, 0, 16)
	// nums = append(nums, 0)
	// nums = append(nums, 1)
	//
	// for len(nums) <= n {
	// 	if nums[len(nums)-1]+nums[len(nums)-2] > mod {
	// 		nums = append(nums, (nums[len(nums)-1]+nums[len(nums)-2])%mod)
	// 	} else {
	// 		nums = append(nums, nums[len(nums)-1]+nums[len(nums)-2])
	// 	}
	// }
	// return nums[len(nums)-1] % mod

	// 方法二: 用滚动数组优化方法一
	// const mod int = 1e9 + 7
	// if n < 2 {
	// 	return n
	// }
	// a, b := 0, 1
	// for i := 2; i <= n; i++ {
	// 	a, b = b, (a+b)%mod
	// }
	// return b

	// 方法三: 矩阵快速幂
	if n < 2 {
		return n
	}
	ret := pow(matrix{{1, 1}, {1, 0}}, n-1)
	return ret[0][0]
}

func main() {
	fmt.Println(fib(2))
	fmt.Println(fib(5))
	fmt.Println(fib(95))
}
