package main

import "sort"

// 多叉树版本 回溯＋剪枝 先排序，每个重复元素的分支只遍历第一个
// 时间O(n*2^n) 空间O(n+logn)=O(n)
//func combinationSum2(candidates []int, target int) (ans [][]int) {
//	sort.Slice(candidates, func(i, j int) bool {
//		return candidates[i] < candidates[j]
//	})
//	var arr []int
//	sum := 0
//	var backTrace func(index int)
//	backTrace = func(index int) {
//		if sum == target {
//			add := make([]int, len(arr))
//			copy(add, arr)
//			ans = append(ans, add)
//			return
//		}
//		for i := index; i < len(candidates); {
//			arr = append(arr, candidates[i])
//			sum += candidates[i]
//			if sum <= target {
//				backTrace(i + 1)
//			}
//			arr = arr[:len(arr)-1]
//			sum -= candidates[i]
//			// 跳过重复的元素
//			for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
//				i++
//			}
//			i++
//		}
//	}
//	backTrace(0)
//	return ans
//}

// 二叉树版本 回溯＋去重 先排序，遍历重复元素的第一个元素时不同层可选取重复的元素, 遍历完第一个重复元素后, 就直接跳过下面重复元素的所有层
// 时间O(n*2^n) 空间O(n+logn)=O(n)
func combinationSum2(candidates []int, target int) (ans [][]int) {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	var arr []int
	sum := 0
	var backTrace func(start, index int)
	backTrace = func(start, index int) {
		if sum == target {
			add := make([]int, len(arr))
			copy(add, arr)
			ans = append(ans, add)
			return
		}
		for index > start && index < len(candidates) && candidates[index-1] == candidates[index] {
			index++
		}
		if index == len(candidates) || sum > target {
			return
		}
		arr = append(arr, candidates[index])
		sum += candidates[index]
		backTrace(index+1, index+1)
		arr = arr[:len(arr)-1]
		sum -= candidates[index]
		backTrace(start, index+1)
	}
	backTrace(0, 0)
	return ans
}
