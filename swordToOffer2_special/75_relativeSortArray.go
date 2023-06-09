package main

// 计数排序 时间O(n+m+upper) 空间O(upper) n为arr1的长度 m为arr2的长度 upper为数组范围长度
func relativeSortArray(arr1 []int, arr2 []int) (ans []int) {
	array := make([]int, 1001)
	for _, num := range arr1 {
		array[num]++
	}
	for _, num := range arr2 {
		for i := 0; i < array[num]; i++ {
			ans = append(ans, num)
		}
		array[num] = 0
	}
	for i, count := range array {
		for j := 0; j < count; j++ {
			ans = append(ans, i)
		}
	}
	return
}
