package main

// 回溯
// 时间O(∑C(k, n)) 空间O(n) ∑C(k, n)为所有组合数集合即解空间, k∈[0, n]
func subsets(nums []int) (ans [][]int) {
	n := len(nums)
	ans = append(ans, []int{})
	var arr []int
	var backtrace func(index int)
	backtrace = func(index int) {
		if index == n {
			return
		}
		arr = append(arr, nums[index])
		add := make([]int, len(arr))
		copy(add, arr)
		ans = append(ans, add)
		backtrace(index + 1)
		arr = arr[:len(arr)-1]
		backtrace(index + 1)
	}
	backtrace(0)
	return
}
