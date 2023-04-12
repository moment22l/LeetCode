package main

import (
	"sort"
)

// 先把二维数组用快排排序再依次遍历看是否能合并 时间O(nlogn+n)=O(nlogn) 空间O(logn)
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var ans [][]int
	now := []int{intervals[0][0], intervals[0][1]}
	for i := 1; i < len(intervals); i++ {
		// 如果当前区间的尾部小于遍历到的区间的首部, 则无法合并, 将当前区间放入ans并使当前区间等于遍历到的区间
		if now[1] < intervals[i][0] {
			add := make([]int, 2)
			copy(add, now)
			ans = append(ans, add)
			now[0] = intervals[i][0]
			now[1] = intervals[i][1]
		}
		// 可以合并的情况就只需要将当前区间的尾部更新为二者尾部的最大值即可
		now[1] = max74(now[1], intervals[i][1])
	}
	add := make([]int, 2)
	copy(add, now)
	ans = append(ans, add)
	return ans
}

func max74(x, y int) int {
	if x > y {
		return x
	}
	return y
}
