package main

func hIndex(citations []int) int {
	// 二分法(官方题解)
	// 答案最多只能到数组长度
	left, right := 0, len(citations)
	var mid int
	for left < right {
		// +1 防止死循环
		mid = (left + right + 1) >> 1
		cnt := 0
		for _, v := range citations {
			if v >= mid {
				cnt++
			}
		}
		if cnt >= mid {
			// 要找的答案在 [mid,right] 区间内
			left = mid
		} else {
			// 要找的答案在 [0,mid) 区间内
			right = mid - 1
		}
	}
	return left
}
