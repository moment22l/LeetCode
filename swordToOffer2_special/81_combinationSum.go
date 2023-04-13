package main

// 时间O(S) 空间O(target) 其中S为所有可行解的长度之和, 递归最差情况就深入到target层
func combinationSum(candidates []int, target int) (ans [][]int) {
	var arr []int
	var backTrace func(sum int, now int)
	backTrace = func(sum int, now int) {
		if sum == target {
			add := make([]int, len(arr))
			copy(add, arr)
			ans = append(ans, add)
			return
		}
		for i := now; i < len(candidates); i++ {
			arr = append(arr, candidates[i])
			sum += candidates[i]
			if sum <= target {
				backTrace(sum, i)
			}
			arr = arr[:len(arr)-1]
			sum -= candidates[i]
		}
	}
	backTrace(0, 0)
	return ans
}
