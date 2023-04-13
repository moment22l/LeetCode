package main

// 回溯 时间O(n^2) 空间O(n+n)=O(n)
//func subsets(nums []int) [][]int {
//	var ans [][]int
//	var now []int
//	var backTrace func(k int)
//	backTrace = func(k int) {
//		if k == len(nums) {
//			add := make([]int, len(now))
//			copy(add, now)
//			ans = append(ans, add)
//			return
//		}
//		now = append(now, nums[k])
//		backTrace(k + 1)
//		now = now[:len(now)-1]
//		backTrace(k + 1)
//	}
//	backTrace(0)
//	return ans
//}

// 迭代 类似二进制 时间O(n*(2^n)) 空间O(n)
func subsets(nums []int) (ans [][]int) {
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		var set []int
		for i, v := range nums {
			if mask>>i&1 > 0 {
				set = append(set, v)
			}
		}
		ans = append(ans, append([]int(nil), set...))
	}
	return
}
