package main

import "sort"

func maxSatisfaction(satisfaction []int) int {
	sort.Slice(satisfaction, func(i, j int) bool {
		return satisfaction[i] > satisfaction[j]
	})
	total := 0
	sum := 0
	for _, num := range satisfaction {
		sum += num
		if sum > 0 {
			total += sum
		} else {
			break
		}
	}
	return total
}
