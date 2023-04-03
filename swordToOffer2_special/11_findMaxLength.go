package main

func findMaxLength(nums []int) (max int) {
	// 用-1代替0即可使用类似题10的方法做
	// 用map记录和为sum的第一个索引(特殊: 前0个元素之和为0)
	sumMap := map[int]int{0: 0}
	sum := 0
	for i := 1; i < len(nums)+1; i++ {
		if nums[i-1] == 0 {
			sum--
		} else {
			sum++
		}
		// 如果sumMap中有与sum相同的key的记录，则i-sumMap[sum]即为01数量一致的长度
		if val, ok := sumMap[sum]; !ok {
			sumMap[sum] = i
		} else {
			if max < i-val {
				max = i - val
			}
		}
	}
	return
}
