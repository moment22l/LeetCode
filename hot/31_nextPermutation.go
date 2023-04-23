package main

// 两遍扫描
func nextPermutation(nums []int) {
	length := len(nums)
	i := length - 2
	// 从后向前找到第一个不是从大到小的数的索引
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		// 如果i不为-1，则寻找i后的序列中第一个比nums[i]大的数
		for j := len(nums) - 1; j > 0; j-- {
			// 找到后将其与i位置交换
			if nums[j] > nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
				break
			}
		}
	}
	reverse(nums[i+1:]) // 逆转i后的从大到小的序列，排序为从小到大的顺序
}

func reverse(nums []int) {
	for i, n := 0, len(nums)-1; i <= n/2; i++ {
		nums[i], nums[n-i] = nums[n-i], nums[i]
	}
}
