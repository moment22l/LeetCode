package main

func singleNumber137(nums []int) int {
	one, two := 0, 0
	for _, num := range nums {
		one = one ^ num&(^two)
		two = two ^ num&(^one)
	}
	return one
}
