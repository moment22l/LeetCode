package main

func tupleSameProduct(nums []int) (ans int) {
	l := len(nums)
	m := make(map[int]int)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			product := nums[i] * nums[j]
			if _, ok := m[product]; ok {
				m[product]++
			} else {
				m[product] = 1
			}
		}
	}
	for _, v := range m {
		ans += (v * (v - 1)) / 2 * 8
	}
	return
}
