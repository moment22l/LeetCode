package main

import "sort"

// 多叉树 回溯＋剪枝 先排序, 剪掉重复元素的子树
// 每次从nums中抽取一个元素出来然后往下递归, 返回后再放回nums 时间O(n*n!), 空间O(n)
func permuteUnique(nums []int) (ans [][]int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
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
			for i+1 < len(nums) && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	backTrace()
	return
}

// 官方题解, 每次填入一个数, 用vis判断是否要遍历当前子树, 如果前面元素重复且vis为false时说明前面重复元素子树已经遍历过, 此时跳过该子树
// 时间O(n*n!), 空间O(n)
//func permuteUnique(nums []int) (ans [][]int) {
//	sort.Ints(nums)
//	n := len(nums)
//	perm := []int{}
//	vis := make([]bool, n)
//	var backtrack func(int)
//	backtrack = func(idx int) {
//		if idx == n {
//			ans = append(ans, append([]int(nil), perm...))
//			return
//		}
//		for i, v := range nums {
//			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
//				continue
//			}
//			perm = append(perm, v)
//			vis[i] = true
//			backtrack(idx + 1)
//			vis[i] = false
//			perm = perm[:len(perm)-1]
//		}
//	}
//	backtrack(0)
//	return
//}
