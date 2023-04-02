package main

import (
	"fmt"
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x := strconv.Itoa(nums[i])
		y := strconv.Itoa(nums[j])
		return x+y < y+x
	})
	res := ""
	for _, num := range nums {
		res += strconv.Itoa(num)
	}
	return res
}

func main() {
	nums := []int{20, 1}
	fmt.Println(minNumber(nums))
}
