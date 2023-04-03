package main

func subarraySum(nums []int, k int) int {
	count, sum := 0, 0
	m := make(map[int]int)
	m[0] = 1
	for _, num := range nums {
		sum += num
		sums, ok := m[sum-k]
		if ok {
			count += sums
		}
		if _, ok = m[sum]; ok {
			m[sum]++
		} else {
			m[sum] = 1
		}
	}
	return count
}
