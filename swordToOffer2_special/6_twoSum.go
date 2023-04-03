package main

//	题目: 排序数组中两个数字之和
//	给定一个已按照 升序排列  的整数数组 numbers ，请你从数组中找出两个数满足相加之和等于目标数 target 。
//	函数应该以长度为 2 的整数数组的形式返回这两个数的下标值。
//	numbers 的下标 从 0 开始计数 ，所以答案数组应当满足 0 <= answer[0] < answer[1] < numbers.length 。
//	假设数组中存在且只存在一对符合条件的数字，同时一个数字不能使用两次。

//func twoSum(numbers []int, target int) []int {
//	// 用一个map记录数和索引的映射，只需每次找map中是否存在target-num就可以知道前面是否含有该数对应的数
//	records := make(map[int]int)
//	for i, num := range numbers {
//		record, ok := records[target-num]
//		if ok {
//			return []int{record, i}
//		}
//		records[num] = i
//	}
//	return []int{}
//}

//func twoSum(numbers []int, target int) []int {
//	// 双指针，p1从头开始，p2从尾开始，小了就右移p1，大了就左移p2
//	i, j := 0, len(numbers)-1
//	for i != j {
//		if numbers[i]+numbers[j] == target {
//			return []int{i, j}
//		} else if numbers[i]+numbers[j] < target {
//			i++
//		} else {
//			j--
//		}
//	}
//	return []int{}
//}

func twoSum(numbers []int, target int) []int {
	// 遍历数组，通过二分查找查询数组中是否有对应数
	for i, num := range numbers {
		j, ok := binarySearch(numbers, target-num)
		if ok && i != j {
			if i < j {
				return []int{i, j}
			} else {
				return []int{j, i}
			}
		}
	}
	return []int{}
}

func binarySearch(numbers []int, target int) (int, bool) {
	left, right := 0, len(numbers)-1
	for left <= right {
		mid := left + (right-left)/2
		if numbers[mid] == target {
			return mid, true
		} else if numbers[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return 0, false
}
