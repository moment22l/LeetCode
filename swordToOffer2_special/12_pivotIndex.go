package main

func pivotIndex(nums []int) int {
	// 笨方法: 记录整个数组的前缀和, 遍历判断当前索引左右数组之和是否一致
	// 以下代码方法: 左右相等时, 当前sum*2+num即为整个数组的和total
	total, sum := 0, 0
	for _, num := range nums {
		total += num
	}
	for i, num := range nums {
		if 2*sum+num == total {
			return i
		}
		sum += num
	}
	return -1
}
