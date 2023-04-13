package main

// 多叉树 每次从nums中抽取一个元素出来然后往下递归, 返回后再放回nums 时间O(n*n!), 空间O(n)
func permute(nums []int) (ans [][]int) {
	var arr []int
	var backTrace func()
	backTrace = func() {
		if len(nums) == 0 {
			add := make([]int, len(arr))
			copy(add, arr)
			ans = append(ans, add)
			return
		}
		for i := 0; i < len(nums); i++ {
			arr = append(arr, nums[i])
			temp := nums[i]
			nums = append(nums[:i], nums[i+1:]...)
			backTrace()
			arr = arr[:len(arr)-1]
			nums = append(nums[:i], append([]int{temp}, nums[i:]...)...)
		}
	}
	backTrace()
	return
}
