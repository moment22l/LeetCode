package main

func numSubarrayProductLessThanK(nums []int, k int) int {
	// 滑动窗口
	left, count, product := 0, 0, 1
	for right := 0; right < len(nums); right++ {
		product *= nums[right]
		for ; left <= right && product >= k; left++ {
			product /= nums[left]
		}
		// 之所以加的是 right-left+1, 是因为每次往右多一个元素, 就会有 right-left+1 个子数组符合条件
		count += right - left + 1
	}
	return count
}
