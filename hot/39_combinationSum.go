package main

// 回溯
func combinationSum(candidates []int, target int) (ans [][]int) {
	arr := make([]int, 0)
	var backtrace func(sum int, i int)
	backtrace = func(sum int, i int) {
		if sum == target {
			add := make([]int, len(arr))
			copy(add, arr)
			ans = append(ans, add)
			return
		}
		if sum > target || i >= len(candidates) {
			return
		}
		sum += candidates[i]
		arr = append(arr, candidates[i])
		backtrace(sum, i)
		sum -= candidates[i]
		arr = arr[:len(arr)-1]
		backtrace(sum, i+1)
	}
	backtrace(0, 0)
	return
}
