package main

import "sort"

func hIndex2(citations []int) int {
	n := len(citations)
	return n - sort.Search(n, func(i int) bool {
		return n-i <= citations[i]
	})
}
