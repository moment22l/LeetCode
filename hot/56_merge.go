package main

import (
	"sort"
)

// 先排序，然后从左到右遍历一遍intervals，没有交叉则直接将now加入到ans，如果有交叉则将当前区间加入到now中
// 时间O(n) 空间O(1)
func merge(intervals [][]int) (ans [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	now := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if now[1] < intervals[i][0] {
			add := make([]int, 2)
			copy(add, now)
			ans = append(ans, now)
			now = intervals[i]
		} else {
			now[1] = max56(intervals[i][1], now[1])
		}
	}
	add := make([]int, 2)
	copy(add, now)
	ans = append(ans, now)
	return
}

func max56(x, y int) int {
	if x > y {
		return x
	}
	return y
}
