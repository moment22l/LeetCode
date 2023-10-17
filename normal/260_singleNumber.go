package main

func singleNumber(nums []int) []int {
	x := 0
	for _, num := range nums {
		x ^= num
	}
	l := x & (-x)
	x1 := 0
	x2 := 0
	for _, num := range nums {
		if num&l == 0 {
			x1 ^= num
		} else {
			x2 ^= num
		}
	}
	return []int{x1, x2}
}
