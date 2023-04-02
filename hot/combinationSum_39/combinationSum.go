package main

import (
	"fmt"
	"sort"
)

var combinations = make([][]int, 0)

func combinationSum(candidates []int, target int) [][]int {
	combinations = [][]int{}
	if len(candidates) == 0 || (len(candidates) == 1 && candidates[0] != target) {
		return [][]int{}
	}
	sort.Ints(candidates)
	index := sort.SearchInts(candidates, target)
	for i := 0; i <= index; i++ {
		combination := make([]int, 0)
		if i < len(candidates) {
			// 问题: 每一次遍历中，如果数组都由当前第一个元素组成，则会产生重复
			for step := 0; step <= index-i; step++ {
				if i+step < len(candidates) {
					backTrack(candidates, target, index, i, step, true, combination)
				}
			}
		}
	}
	return combinations
}

func backTrack(candidates []int, target int, index int, now int, step int, flag bool, combination []int) {
	total := sum(combination)
	if total == target {
		copySlice := copy(combination)
		combinations = append(combinations, copySlice)
	} else {
		if total < target && now <= index && now < len(candidates) {
			temp := now
			for i := 0; i <= (target-total)/candidates[temp]; i++ {
				combination = append(combination, candidates[temp])
				if flag == true {
					now = now + step
					flag = false
				}
				backTrack(candidates, target, index, now+1, step, flag, combination)
			}
		}
	}
}

func sum(nums []int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total = total + nums[i]
	}
	return total
}

func copy(nums []int) []int {
	copySlice := make([]int, 0)
	for _, num := range nums {
		copySlice = append(copySlice, num)
	}
	return copySlice
}

func main() {
	//candidates := []int{2, 3, 6, 7}
	//target := 7
	//fmt.Println(combinationSum(candidates, target))
	//candidates = []int{2, 3, 5}
	//target = 8
	//fmt.Println(combinationSum(candidates, target))
	//candidates = []int{2}
	//target = 1
	//fmt.Println(combinationSum(candidates, target))
	//candidates = []int{3, 5, 8}
	//target = 11
	//fmt.Println(combinationSum(candidates, target))
	candidates := []int{7, 3, 2}
	target := 18
	fmt.Println(combinationSum(candidates, target))
	//candidates := []int{3, 2}
	//target := 6
	//fmt.Println(combinationSum(candidates, target))
}
