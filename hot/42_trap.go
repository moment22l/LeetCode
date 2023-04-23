package main

// 动态规划 时间O(n) 空间O(n)
// 先记录每一列的左右两边最高的墙的高度
// 遍历每一列, 每一列中可存储的容量便为左右两边最高的墙里较矮的墙与本列墙高度的差值(较矮墙比本列高时)
//func trap(height []int) (ans int) {
//	maxLeft := make([]int, len(height))
//	maxRight := make([]int, len(height))
//	for i := 1; i < len(maxLeft)-1; i++ {
//		maxLeft[i] = max42(maxLeft[i-1], height[i-1])
//	}
//	for i := len(maxRight) - 2; i > 0; i-- {
//		maxRight[i] = max42(maxRight[i+1], height[i+1])
//	}
//	for i := 1; i < len(height)-1; i++ {
//		if min42(maxLeft[i], maxRight[i]) > height[i] {
//			ans += min42(maxLeft[i], maxRight[i]) - height[i]
//		}
//	}
//	return
//}

// 双指针, 当height[left] < height[right]时右移left, 反之左移right
// 时间O(n) 空间O(1)
func trap(height []int) (ans int) {
	left, right := 0, len(height)-1
	maxLeft, maxRight := 0, 0
	for left < right {
		maxLeft = max42(maxLeft, height[left])
		maxRight = max42(maxRight, height[right])
		if height[left] < height[right] {
			ans += maxLeft - height[left]
			left++
		} else {
			ans += maxRight - height[right]
			right--
		}
	}
	return
}

func max42(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min42(x, y int) int {
	if x < y {
		return x
	}
	return y
}
