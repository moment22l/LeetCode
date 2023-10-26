package main

import "sort"

func avoidFlood(rains []int) []int {
	ans := make([]int, len(rains))
	// 将所有日子抽的都先置1，后续更新（返回时依旧没有使用的湖泊会默认为1
	for i := range ans {
		ans[i] = 1
	}
	m := make(map[int]int) // key=湖泊id  val=装满该湖泊的日子
	q := make([]int, 0)    // 队列, 记录了没有下雨且没有被使用到的索引
	for today, lakeID := range rains {
		// 判断是抽水还是放水
		if lakeID == 0 {
			// 如果是抽水，那就把当前索引放入到q中
			q = append(q, today)
			continue
		}
		// 将对应天数的ans置为-1
		ans[today] = -1
		// 如果是放水，那就判断将要满上的湖泊是否已经满了
		if fullDay, ok := m[lakeID]; ok {
			// 利用sort中的SearchInts函数找到满足(晚于即大于fullDay)的元素索引
			index := sort.SearchInts(q, fullDay)
			// 没找到就直接返回空数组
			if index == len(q) {
				return []int{}
			}
			// 如果有可以满足条件的元素，就更新ans数组
			ans[q[index]] = lakeID
			q = append(q[:index], q[index+1:]...)
		}
		// 更新对应湖泊最后满上的日子
		m[lakeID] = today
	}
	return ans
}
