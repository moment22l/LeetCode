package main

import "LeetCode/utils"

// 最简单的方法，两两合并
// 时间O(n*m^2) 空间O(1), m=len(lists), n=max{len(lists[i])}, i∈[0, len(lists)-1]
func mergeKLists(lists []*utils.ListNode) *utils.ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) != 1 {
		lists[1] = mergeTwoLists(lists[0], lists[1])
		lists = lists[1:]
	}
	return lists[0]
}
