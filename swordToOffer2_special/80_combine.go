package main

// 回溯 时间O(C(n,k)*k) 空间O(k)  ps: C(n,k)表示n个不同元素中取出k个元素的组合数
func combine(n int, k int) [][]int {
	var ans [][]int
	var arr []int
	var backTrace func(now int)
	backTrace = func(now int) {
		if len(arr)+(n-now+1) < k {
			// 用于剪枝, 如果当前数组长度加上数列后面所有数的长度比k小, 就不可能再构造出子集了
			return
		}
		if len(arr) == k {
			add := make([]int, k)
			copy(add, arr)
			ans = append(ans, add)
			return
		}
		arr = append(arr, now)
		backTrace(now + 1)
		arr = arr[:len(arr)-1]
		backTrace(now + 1)
	}
	backTrace(1)
	return ans
}
