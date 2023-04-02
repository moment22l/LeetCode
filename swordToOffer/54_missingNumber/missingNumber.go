package main

func missingNumber(nums []int) int {
	left, right := 0, len(nums)-1
	var middle int
	for left <= right {
		middle = (left + right) / 2
		if nums[middle] == middle {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}

func main() {

}
