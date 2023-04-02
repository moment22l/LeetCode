package main

import "math"

//	题目: 整数除法
//	给定两个整数 a 和 b ，求它们的除法的商 a/b ，要求不得使用乘号 '*'、除号 '/' 以及求余符号 '%' 。
//	注意：
//	整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2
//	假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231, 231−1]。本题中，如果除法结果溢出，则返回 231 − 1

// 利用快速幂算法的思想算快速乘
func quickAdd(z, y, x int) bool {
	for result, add := 0, y; z > 0; z >>= 1 {
		if z&1 > 0 {
			if result < x-add {
				return false
			}
			result += add
		}
		if z != 1 {
			if add < x-add {
				return false
			}
			add += add
		}
	}
	return true
}

func divide(a int, b int) int {
	// 边界情况
	if a == math.MinInt32 {
		if b == 1 {
			return math.MinInt32
		} else if b == -1 {
			return math.MaxInt32
		}
	} else if a == 0 {
		return 0
	}
	if b == math.MinInt32 {
		if a == math.MinInt32 {
			return 1
		}
		return 0
	}
	rev := false
	if a > 0 {
		a = -a
		rev = !rev
	}
	if b > 0 {
		b = -b
		rev = !rev
	}
	ans := 0

	// 二分查找
	//left, right := 1, math.MaxInt32
	//for left <= right {
	//	mid := left + (right-left)>>1
	//	if quickAdd(mid, b, a) {
	//		ans = mid
	//		if mid == math.MaxInt32 {
	//			break
	//		}
	//		left = mid + 1
	//	} else {
	//		right = mid - 1
	//	}
	//}

	// 类二分查找
	candidates := []int{b}
	for y := b; y >= a-y; {
		y += y
		candidates = append(candidates, y)
	}
	for i := len(candidates) - 1; i >= 0; i-- {
		if candidates[i] >= a {
			ans |= 1 << i
			a -= candidates[i]
		}
	}
	if rev {
		return -ans
	} else {
		return ans
	}
}
