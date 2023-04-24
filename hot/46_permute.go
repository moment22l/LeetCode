package main

// 回溯 原地改变nums(删除选出的元素)
// 时间O(n*n!) 空间O(n)
//func permute(nums []int) (ans [][]int) {
//	arr := make([]int, 0)
//	var backtrace func()
//	backtrace = func() {
//		if len(nums) == 0 {
//			add := make([]int, len(arr))
//			copy(add, arr)
//			ans = append(ans, add)
//			return
//		}
//		for i := 0; i < len(nums); i++ {
//			temp := nums[i]
//			arr = append(arr, nums[i])
//			nums = append(nums[:i], nums[i+1:]...)
//			backtrace()
//			nums = append(nums[:i], append([]int{temp}, nums[i:]...)...)
//			arr = arr[:len(arr)-1]
//		}
//	}
//	backtrace()
//	return
//}

func permute(nums []int) (ans [][]int) {
	var backtrace func(int)
	backtrace = func(first int) {
		if first == len(nums) {
			add := make([]int, len(nums))
			copy(add, nums)
			ans = append(ans, add)
			return
		}
		for i := first; i < len(nums); i++ {
			nums[i], nums[first] = nums[first], nums[i]
			backtrace(first + 1)
			nums[i], nums[first] = nums[first], nums[i]
		}
	}
	backtrace(0)
	return
}
