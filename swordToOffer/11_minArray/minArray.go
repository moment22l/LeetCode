package main

func minArray(numbers []int) int {
	// 二分法
	left := 0
	right := len(numbers) - 1
	middle := left + (right-left)/2
	for left < right {
		if numbers[middle] > numbers[right] {
			// middle大于right说明最小值在middle-right之间，且不可能为middle
			left = middle + 1
		} else if numbers[middle] < numbers[right] {
			// middle小于right说明最小值在left-middle之间，且可能为middle
			right = middle
		} else {
			// middle等于right说明最小值两边都有可能，但由于有middle作为right的替代，所以可以排除掉right
			right--
		}
		middle = left + (right-left)/2
	}
	return numbers[middle]
}

func main() {

}
