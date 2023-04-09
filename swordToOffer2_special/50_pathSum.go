package main

import "LeetCode/utils"

// 时间O(n^2)(遍历n个节点同时每个节点都要遍历一次total数组) 空间O(n)
func pathSum(root *utils.TreeNode, targetSum int) int {
	count := 0
	total := make([]int, 0)
	var dfs func(*utils.TreeNode)
	dfs = func(root *utils.TreeNode) {
		if root == nil {
			return
		}
		total = append(total, 0)
		arrayAdd(total, root.Val)
		num := checkArray(total, targetSum)
		count += num
		dfs(root.Left)
		dfs(root.Right)
		arrayAdd(total, (-1)*root.Val)
		total = total[:len(total)-1]
	}
	dfs(root)
	return count
}

func arrayAdd(arr []int, k int) []int {
	for i := range arr {
		arr[i] += k
	}
	return arr
}

func checkArray(arr []int, k int) int {
	count := 0
	for _, num := range arr {
		if num == k {
			count++
		}
	}
	return count
}

// 前缀和 时间O(n) 空间O(n)
//func pathSum(root *utils.TreeNode, targetSum int) int {
//	count := 0
//	preSum := map[int64]int{0: 1}
//	var dfs func(*utils.TreeNode, int64)
//	dfs = func(root *utils.TreeNode, cur int64) {
//		if root == nil {
//			return
//		}
//		cur += int64(root.Val)
//		count += preSum[cur-int64(targetSum)]
//		preSum[cur]++
//		dfs(root.Left, cur)
//		dfs(root.Right, cur)
//		preSum[cur]--
//	}
//	dfs(root, 0)
//	return count
//}
