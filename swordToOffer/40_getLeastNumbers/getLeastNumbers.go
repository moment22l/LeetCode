package main

import (
	"fmt"
	"sort"
)

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	if k == len(arr) {
		return arr
	}
	sort.Ints(arr)
	return arr[:k]
	//return quickSort(arr, 0, len(arr)-1, k)
}

func quickSort(arr []int, left int, right int, k int) []int {
	pivot := arr[left]
	i, j := left, right
	for i < j {
		for i < j && arr[j] >= pivot {
			j--
		}
		for i < j && arr[i] <= pivot {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[i], arr[left] = arr[left], arr[i]
	if i < k {
		return quickSort(arr, i+1, right, k)
	} else if i > k {
		return quickSort(arr, left, i-1, k)
	} else {
		return arr[:i]
	}
}

func main() {
	fmt.Println(getLeastNumbers([]int{1, 0, 5, 4, 2, 4, 3}, 4))
}
