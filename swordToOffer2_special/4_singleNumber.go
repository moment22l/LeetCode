package main

//	题目: 只出现一次的数字
//	给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。

func singleNumber(nums []int) int {
	bits := make([]int, 32)
	for _, num := range nums {
		for i := 0; i < 32; i++ {
			bits[i] += (num >> (31 - i)) & 1
		}
	}

	var result int
	for i := 0; i < 32; i++ {
		result = (result << 1) + bits[i]%3
	}
	result32 := int32(result)
	result = int(result32)
	return result
}
